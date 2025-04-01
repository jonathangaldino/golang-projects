package shipping

import "go-event-driven/internal/eventbus"

type ShippingService struct {
	eventBus *eventbus.EventBus
}

func NewShippingService(eventBus *eventbus.EventBus) *ShippingService {
	return &ShippingService{
		eventBus: eventBus,
	}
}
