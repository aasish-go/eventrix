## Phased Plan — Eventrix v0.1.0 MVP

We are building Eventrix MVP in clear phases → to ensure quality and encourage contribution:

---

### Phase 1 — Project Foundation (Week 1)

- Project Setup
    - Setup GitHub repo and initial documentation
    - Create Project Board and Issues
- Core API + InMemoryBus
    - Define EventBus and EventHandler interfaces
    - Implement InMemoryBus driver
    - Write unit tests and example app

---

### Phase 2 — KafkaBus Integration (Week 2)

- Implement KafkaBus publish
- Implement KafkaBus subscribe
- Write tests for KafkaBus

---

### Phase 3 — Processor: SubscriptionRouter + Retry + DLQ (Week 3)

- Implement SubscriptionRouter with Retry Support
- Implement DLQProcessor
- Write tests for DLQ

---

### Phase 4 — Observability (Week 4)

- Integrate OpenTelemetry tracing
- Example app with tracing
- Implement Prometheus metrics
- Test metrics endpoint

---

### Phase 5 — Final Polish & v0.1.0 Release (Week 5)

- Polish Usage Guide and examples
- Add GitHub Actions CI + Makefile + PR templates
- Polish repo and tag v0.1.0 release

---

### Future Phases (Post MVP)

- NATSBus driver support
- RedisBus driver support
- Multi-transport SubscriptionRouter
- CLI tooling (Eventrix CLI)
- Pluggable retry policies
- Event replay support
- Tracing visualizer
- Schema registry integration

---

## How to contribute to current phase

- Check our [GitHub Project Board](https://github.com/YOUR_REPO_LINK/projects) → see which phase is active
- Pick Stories from the current phase → open PRs → follow our [CONTRIBUTING.md](CONTRIBUTING.md)

---

 Let’s build Eventrix together!
