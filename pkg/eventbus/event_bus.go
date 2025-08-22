package eventbus

import (
	"context"

	"github.com/aasish-go/eventrix/pkg/eventhandler"
)

// EventBus defines the Publish/Subscribe contract for Eventrix.
type EventBus interface {
	Publish(ctx context.Context, topic string, event interface{}) error
	Subscribe(ctx context.Context, topic string, handler eventhandler.EventHandler) error
	Close() error
}
