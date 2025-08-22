package bus

import (
	"context"
	"errors"
	"sync"

	"github.com/aasish-go/eventrix/pkg/eventhandler"
)

type InMemoryBus struct {
	mu       sync.RWMutex
	handlers map[string][]eventhandler.EventHandler
}

func NewInMemoryBus() *InMemoryBus {
	return &InMemoryBus{
		handlers: make(map[string][]eventhandler.EventHandler),
	}
}

func (b *InMemoryBus) Publish(ctx context.Context, topic string, event interface{}) error {
	b.mu.RLock()
	defer b.mu.RUnlock()

	hs, ok := b.handlers[topic]
	if !ok {
		return errors.New("no handlers for topic: " + topic)
	}

	for _, handler := range hs {
		if err := handler.HandleEvent(ctx, event); err != nil {
			// You can later add retry logic here
			return err
		}
	}

	return nil
}

func (b *InMemoryBus) Subscribe(ctx context.Context, topic string, handler eventhandler.EventHandler) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.handlers[topic] = append(b.handlers[topic], handler)
	return nil
}

func (b *InMemoryBus) Close() error {
	// No-op for InMemoryBus
	return nil
}
