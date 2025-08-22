package eventhandler

import "context"

// EventHandler defines the handler interface for processing events.
type EventHandler interface {
	HandleEvent(ctx context.Context, event interface{}) error
}
