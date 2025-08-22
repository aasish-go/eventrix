# Eventrix — Usage Guide
### Overview
This guide shows how to use the `Eventrix` framework to:

- Publish events
- Subscribe to events
- Use different bus drivers (Kafka, NATS, InMemory)
- Configure tracing and metrics
- Configure retries and DLQ

### Setup
#### Installation
```gotemplate
go get github.com/aasish-go/eventrix
```
### Basic Example
#### Publish an Event
```gotemplate
orderEvent := OrderPlacedEvent{
OrderID: "12345",
Amount:  99.99,
}

err := bus.Publish(ctx, "orders.placed", orderEvent)
if err != nil {
log.Fatalf("Failed to publish event: %v", err)
}
```

#### Subscribe to an Event
```gotemplate
err := bus.Subscribe(ctx, "orders.placed", func(ctx context.Context, event Event) error {
order := event.Payload().(OrderPlacedEvent)
log.Printf("Processing order %s", order.OrderID)
return nil
})

if err != nil {
log.Fatalf("Failed to subscribe: %v", err)
}
```

### Supported Bus Drivers
| Driver    |             How to configure              |
|:----------|:-----------------------------------------:|
| Kafka     |          See config.yaml example          |
| NATS      |          See config.yaml example          |
| InMemory  | No config needed → useful for unit tests  |

### Configuration Example
See [docs/config.yaml](config.yaml).

Key parameters:
- Kafka brokers list
- NATS server URL
- Retry policies (max retries, backoff)
- DLQ topic
- Tracing options
- Metrics options

### Observability
#### Tracing
The framework automatically creates OpenTelemetry spans for:
- Publish
- Consume
- Event handler execution

#### Metrics
Prometheus metrics exposed at /metrics endpoint:
- events_published_total
- events_consumed_total
- event_handler_latency_seconds
- events_retried_total
- events_dlq_total

### Testing
Use InMemoryBus for unit and integration tests:
```gotemplate
bus := framework.NewInMemoryBus()
// Use bus.Publish and bus.Subscribe as usual
```



