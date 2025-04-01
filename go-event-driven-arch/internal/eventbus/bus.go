package eventbus

import (
	"sync"
)

type EventBus struct {
	subscribers map[string][]chan Event
	mu          sync.RWMutex
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]chan Event),
	}
}

func (b *EventBus) Subscribe(eventType string, subscriber chan Event) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.subscribers[eventType] = append(b.subscribers[eventType], subscriber)
}

func (b *EventBus) Publish(event Event) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	subscribers, exists := b.subscribers[event.Type]

	if !exists {
		return
	}

	for _, ch := range subscribers {
		go func(ch chan Event) {
			ch <- event
		}(ch)
	}
}
