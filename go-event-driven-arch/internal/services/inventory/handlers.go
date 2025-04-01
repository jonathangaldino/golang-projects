package inventory

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go-event-driven/internal/domain"
	"go-event-driven/internal/eventbus"
	"log"
	"time"
)

func (s *InventoryService) Start(ctx context.Context) {
	ch := make(chan eventbus.Event, 10)
	s.eventBus.Subscribe("order.created", ch)

	go func() {
		for {
			select {
			case event := <-ch:
				s.handleOrderCreated(event)

			case <-ctx.Done():
				return
			}

		}
	}()
}

var ErrInsufficientStock = errors.New("insufficient stock available")

func (s *InventoryService) handleOrderCreated(event eventbus.Event) {
	var order domain.Order

	if err := json.Unmarshal(event.Payload, &order); err != nil {
		log.Printf("Error unmarshalling order: %v", err)
		return
	}

	if err := s.processOrder(order); err != nil {
		s.publishOrderFailedEvent(order, err.Error())
		return
	}

	s.publishInventoryUpdatedEvent(order)
}

func (s *InventoryService) processOrder(order domain.Order) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	insufficientItems := make([]string, 0)

	for _, product := range order.Products {
		currentStock, exists := s.inventory[product.ID]

		if !exists {
			insufficientItems = append(insufficientItems, fmt.Sprintf("%s (not found)", product.ID))
			continue
		}

		if currentStock < product.Quantity {
			insufficientItems = append(insufficientItems, fmt.Sprintf("%s (request: %d, available: %d)", product.ID, product.Quantity, currentStock))
		}
	}

	if len(insufficientItems) > 0 {
		errorDetails := fmt.Sprintf("Insufficient inventory for items: %v", insufficientItems)
		log.Printf("Order %s failed: %s", order.ID, errorDetails)
		return fmt.Errorf("%w: %s", ErrInsufficientStock, errorDetails)
	}

	for _, product := range order.Products {
		currentStock := s.inventory[product.ID]
		newStock := currentStock - product.Quantity
		s.inventory[product.ID] = newStock

		fmt.Printf("Inventory updated for product %s: %d -> %d\n", product.ID, currentStock, newStock)

		if newStock < 10 {
			s.publishLowStockEvent(product.ID, newStock)
		}
	}

	return nil
}

func (s *InventoryService) publishLowStockEvent(productID string, newStock int) {
	lowStockProduct := map[string]interface{}{
		"product_id": productID,
		"stock":      newStock,
	}

	payload, _ := json.Marshal(lowStockProduct)
	s.eventBus.Publish(eventbus.Event{
		Type:      "inventory.low_stock",
		Timestamp: time.Now(),
		Payload:   payload,
	})
}

func (s *InventoryService) publishInventoryUpdatedEvent(order domain.Order) {
	updatedInventory := map[string]interface{}{
		"order_id":  order.ID,
		"success":   true,
		"timestamp": time.Now(),
	}

	payload, err := json.Marshal(updatedInventory)

	if err != nil {
		log.Printf("Error marshalling inventory updated event: %v", err)
		return
	}

	s.eventBus.Publish(eventbus.Event{
		Type:      domain.EventInventoryUpdated,
		Timestamp: time.Now(),
		Payload:   payload,
	})
}

func (s *InventoryService) publishOrderFailedEvent(order domain.Order, reason string) {
	failedOrder := map[string]interface{}{
		"order_id":  order.ID,
		"reason":    reason,
		"timestamp": time.Now(),
	}

	payload, err := json.Marshal(failedOrder)

	if err != nil {
		log.Printf("Error marshalling order failed event: %v", err)
		return
	}

	s.eventBus.Publish(eventbus.Event{
		Type:      domain.EventOrderFailed,
		Timestamp: time.Now(),
		Payload:   payload,
	})

	log.Printf("Published order.failed event for order %s", order.ID)
}
