package notification

import (
	"context"
	"encoding/json"
	"fmt"
	"go-event-driven/internal/domain"
	"go-event-driven/internal/eventbus"
)

func (s *NotificationService) Start(ctx context.Context) {
	orderCh := make(chan eventbus.Event, 10)
	s.eventBus.Subscribe("order.created", orderCh)

	lowStockCh := make(chan eventbus.Event, 10)
	s.eventBus.Subscribe("inventory.low_stock", lowStockCh)

	go func() {
		for {
			select {
			case event := <-orderCh:
				s.handleOrderCreated(event)

			case event := <-lowStockCh:
				s.handleInventoryLowStock(event)

			case <-ctx.Done():
				return
			}
		}
	}()
}

func (s *NotificationService) handleOrderCreated(event eventbus.Event) {
	var order domain.Order

	if err := json.Unmarshal(event.Payload, &order); err != nil {
		return
	}

	fmt.Printf("Notification sent to customer %s: Order %s received! \n", order.CustomerID, order.ID)
}

func (s *NotificationService) handleInventoryLowStock(event eventbus.Event) {
	var data map[string]interface{}

	if err := json.Unmarshal(event.Payload, &data); err != nil {
		return
	}

	fmt.Printf("Alert: Low stock for product %v (remaining: %v)\n", data["product_id"], data["stock"])

}
