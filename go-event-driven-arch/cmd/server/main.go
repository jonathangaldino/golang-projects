package main

import (
	"context"
	"fmt"
	"go-event-driven/internal/domain"
	"go-event-driven/internal/eventbus"
	"go-event-driven/internal/services/inventory"
	"go-event-driven/internal/services/notification"
	"go-event-driven/internal/services/order"
	"go-event-driven/internal/services/shipping"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	bus := eventbus.NewEventBus()

	// Initialize services
	orderSvc := order.NewService(bus)
	inventorySvc := inventory.NewInventoryService(bus)
	notificationSvc := notification.NewNotificationService(bus)
	shippingSvc := shipping.NewShippingService(bus)

	// Create context that can be canceled
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	inventorySvc.Start(ctx)
	notificationSvc.Start(ctx)
	shippingSvc.Start(ctx)

	// Create some sample orders

	order1 := domain.Order{
		ID:         "order123",
		CustomerID: "customer123",
		Products: []domain.Product{
			{
				ID:       "p1",
				Name:     "Product 1",
				Quantity: 10,
			},
			{
				ID:       "p2",
				Name:     "Product 2",
				Quantity: 5,
			},
		},
		Status: "created",
	}

	order2 := domain.Order{
		ID:         "order123",
		CustomerID: "customer123",
		Products: []domain.Product{
			{
				ID:       "p3",
				Name:     "Product 3",
				Quantity: 2,
			},
		},
		Status: "created",
	}

	fmt.Println("Creating orders...")
	orderSvc.CreateOrder(order1)

	time.Sleep(1 * time.Second)
	fmt.Println("\n-------------")

	orderSvc.CreateOrder(order2)

	// Wait for termination signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down services...")
	cancel()
	time.Sleep(1 * time.Second)
}
