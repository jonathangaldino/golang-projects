package order

import (
	"encoding/json"
	"fmt"
	"go-event-driven/internal/domain"
	"go-event-driven/internal/eventbus"
	"sync"
	"time"
)

type Service struct {
	eventBus *eventbus.EventBus
	orders   map[string]domain.Order
	mu       sync.RWMutex
}

func NewService(eventBus *eventbus.EventBus) *Service {
	return &Service{
		eventBus: eventBus,
		orders:   make(map[string]domain.Order),
	}
}

func (s *Service) CreateOrder(order domain.Order) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.orders[order.ID] = order

	orderJSON, err := json.Marshal(order)

	if err != nil {
		return err
	}

	s.eventBus.Publish(eventbus.Event{
		Type:      "order.created",
		Timestamp: time.Now(),
		Payload:   orderJSON,
	})

	fmt.Printf("Order created: %s\n", order.ID)
	return nil
}
