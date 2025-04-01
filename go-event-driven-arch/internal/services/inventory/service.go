package inventory

import (
	"go-event-driven/internal/eventbus"
	"sync"
)

type InventoryService struct {
	eventBus  *eventbus.EventBus
	inventory map[string]int
	mu        sync.RWMutex
}

func NewInventoryService(eventBus *eventbus.EventBus) *InventoryService {
	service := &InventoryService{
		eventBus:  eventBus,
		inventory: make(map[string]int),
	}

	// Initialize inventory with some products.
	service.inventory["p1"] = 10
	service.inventory["p2"] = 25
	service.inventory["p3"] = 0

	return service
}
