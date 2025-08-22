# Eventrix — Event-driven Framework for Go 🚀
<img src="/images/logo.png" alt="Alt text" width="250" height="250">

**Eventrix** is a modern, pluggable, high-performance event-driven framework written in Go.  
It provides a clean abstraction for building **event-driven services** with support for multiple message buses and powerful observability.

## ✨ Key Features

- ✅ Pluggable transport drivers → Kafka, InMemory, NATS (planned), Redis (planned)
- ✅ SubscriptionRouter with integrated Retry and DLQ handling
- ✅ First-class observability → OpenTelemetry Tracing, Prometheus Metrics
- ✅ Simple core API → easy to embed into any Go service
- ✅ Production-ready architecture → based on battle-tested patterns
- ✅ Modular → CLI and tooling planned

## 🚀 Roadmap

### 🎯 v0.1.0 MVP (June-July 2025)

- [x] Project Setup
- [ ] Core API + InMemoryBus
- [ ] KafkaBus
- [ ] Processor (SubscriptionRouter + Retry + DLQ)
- [ ] Observability
- [ ] Final Polish & Release

### 🎯 v0.2.0 (Planned)

- NATSBus driver support
- RedisBus driver support
- Multi-transport support
- CLI tooling (Eventrix CLI)
- Improved retry policies (pluggable)

### 🎯 v1.0.0 (Stable Release)

- Production benchmarks & tuning
- Comprehensive documentation
- Operator patterns & cookbook
- Plugin architecture
- Helm chart for Kubernetes
- First production adoption

## 👥 Contributing

We welcome contributions!  
Please read [CONTRIBUTING.md](docs/CONTRIBUTING.md) for guidelines.

## 📄 License

MIT License — see [LICENSE](LICENSE).
