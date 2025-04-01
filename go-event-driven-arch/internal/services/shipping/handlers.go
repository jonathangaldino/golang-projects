package shipping

import (
	"context"
	"encoding/json"
	"fmt"
	"go-event-driven/internal/domain"
	"go-event-driven/internal/eventbus"
	"time"
)

func (s *ShippingService) Start(ctx context.Context) {
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

func (s *ShippingService) handleOrderCreated(event eventbus.Event) {
	var order domain.Order

	if err := json.Unmarshal(event.Payload, &order); err != nil {
		return
	}

	// Simulate shipping process
	time.Sleep(2 * time.Second)

	fmt.Printf("Order %s prepared for shipping\n", order.ID)

	order.Status = "shipped"
	orderJSON, _ := json.Marshal(order)

	s.eventBus.Publish(eventbus.Event{
		Type:      "order.shipped",
		Timestamp: time.Now(),
		Payload:   orderJSON,
	})
}
