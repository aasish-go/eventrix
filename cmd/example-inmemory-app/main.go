package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aasish-go/eventrix/internal/bus"
)

type UserCreatedEvent struct {
	UserID string
	Email  string
}

type UserCreatedHandler struct{}

func (h *UserCreatedHandler) HandleEvent(ctx context.Context, event interface{}) error {
	evt, ok := event.(UserCreatedEvent)
	if !ok {
		return fmt.Errorf("invalid event type")
	}

	log.Printf("Handled UserCreatedEvent: %+v\n", evt)
	return nil
}

func main() {
	ctx := context.Background()

	bus := bus.NewInMemoryBus()

	handler := &UserCreatedHandler{}
	_ = bus.Subscribe(ctx, "user.created", handler)

	_ = bus.Publish(ctx, "user.created", UserCreatedEvent{
		UserID: "123",
		Email:  "test@example.com",
	})
	_ = bus.Publish(ctx, "user.created", UserCreatedEvent{
		UserID: "123",
		Email:  "test@example.com",
	})

	_ = bus.Close()
}
