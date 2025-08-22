
---

# Ready-to-copy Folder Structure section for your README.md or ARCHITECTURE.md

```markdown
Project Folder Structure
```
```plaintext
eventrix/                        # Root of the repo
├── cmd/                         # Example apps, CLI tools
│   ├── example-inmemory-app/    # Example app using InMemoryBus
│   ├── example-kafka-app/       # Example app using KafkaBus
│   ├── example-tracing-app/     # Example app demonstrating tracing
│   └── example-observability-app/ # Example app with Prometheus metrics
│
├── pkg/                         # Public framework API (stable for users)
│   ├── eventbus/                # EventBus interface and abstractions
│   ├── eventhandler/            # EventHandler interface and abstractions
│   ├── event/                   # Event definitions and helpers
│
├── internal/                    # Internal implementation (subject to change, not public API)
│   ├── bus/                     # Bus driver implementations
│   │   ├── inmemory_bus.go
│   │   ├── kafka_bus.go
│   │   ├── nats_bus.go          # (planned)
│   │   ├── redis_bus.go         # (planned)
│   │   └── bus_utils.go
│   ├── processor/               # SubscriptionRouter, RetryProcessor, DLQProcessor
│   │   ├── subscription_router.go
│   │   ├── retry_processor.go
│   │   ├── dlq_processor.go
│   │   └── processor_utils.go
│   ├── tracing/                 # OpenTelemetry Tracing integration
│   │   └── otel_tracer.go
│   ├── metrics/                 # Prometheus metrics collector
│   │   └── metrics_collector.go
│
├── configs/                     # Example configs (config.yaml, docker-compose.yaml, etc)
│   ├── config.yaml
│   ├── docker-compose-kafka.yaml
│   ├── docker-compose-observability.yaml
│
├── docs/                        # Project documentation
│   ├── ARCHITECTURE.md
│   ├── USAGE_GUIDE.md
│   ├── CONTRIBUTING.md
│   ├── GIT_GUIDELINES.md
│   ├── ROADMAP.md
│   ├── PROJECT_PLAN.md
│
├── internal/testing/            # Unit tests and integration tests
│   ├── inmemory_bus_test.go
│   ├── kafka_bus_test.go
│   ├── dlq_processor_test.go
│   ├── subscription_router_test.go
│   └── metrics_test.go
│
├── .github/                     # GitHub settings
│   ├── ISSUE_TEMPLATE/          # Issue templates
│   ├── PULL_REQUEST_TEMPLATE.md # PR template
│   └── workflows/               # GitHub Actions CI workflows
│       └── ci.yml
│
├── go.mod
├── go.sum
├── Makefile                     # Build / test / lint targets
├── LICENSE
└── README.md
```