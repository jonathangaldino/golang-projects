package notification

import "go-event-driven/internal/eventbus"

type NotificationService struct {
	eventBus *eventbus.EventBus
}

func NewNotificationService(eventBus *eventbus.EventBus) *NotificationService {
	return &NotificationService{
		eventBus: eventBus,
	}
}
