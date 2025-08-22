package event

// Event is the base abstraction for all events in Eventrix.
type Event struct {
	Topic string
	Data  interface{}
}
