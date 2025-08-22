# Eventrix — Architecture Design

## Overview

Eventrix is a modular, pluggable **event-driven framework for Go**, designed for building reliable event-driven systems with first-class observability.

It is structured to provide:

* Clean abstractions → EventBus, EventHandler  
* Pluggable drivers → Kafka, InMemory, NATS (planned), Redis (planned)  
* SubscriptionRouter → automatic retry, DLQ support  
* Observability → OpenTelemetry Tracing, Prometheus Metrics

---

## High-Level Architecture
```plaintext
+-------------+ +--------------------+ +-------------------+
| Producer | --Publish->| EventBus |---> | SubscriptionRouter |
+-------------+ +--------------------+ +-------------------+
|
V
+-------------+
| EventHandler |
+-------------+
|
V
+---------------+
| RetryProcessor |
+---------------+
|
V
+------------+
| DLQProcessor |
+------------+
```
---

## Core Components

### EventBus

Abstract interface for **publish/subscribe**.
- Publish(Event) → sends event to bus
- Subscribe(EventHandler) → registers handler

Drivers:
- InMemoryBus → for testing, dev
- KafkaBus → production-ready
- NATSBus → planned
- RedisBus → planned

---

### SubscriptionRouter

Dispatches events to correct EventHandler(s).
- Manages handler registration
- Supports automatic **RetryProcessor** and **DLQProcessor** integration

---

### RetryProcessor

Wraps EventHandler with retry logic:
- Configurable max retries
- Exponential backoff
- Failing events after retries → sent to DLQProcessor

---

### DLQProcessor

Handles permanently failed events:
- Publishes to DLQ topic
- Allows operators to inspect dead-letter events

---

### Observability

- OpenTelemetry Tracing → spans for Publish, Subscribe, Handler execution
- Prometheus Metrics → /metrics endpoint with:
    - Events published
    - Events consumed
    - Handler latency
    - Retry count
    - DLQ count

---

## Planned Extensions

- Multi-transport SubscriptionRouter → route event to correct transport
- Event replay support
- CLI tooling → event monitor, replay CLI
- Schema registry integration

---

## Summary

Eventrix aims to provide:
- A production-quality event-driven framework for Go
- Minimal API → easy to adopt
- First-class observability out of the box
- Pluggable architecture → grow with your system

**MVP → v0.1.0 → implements this core architecture.**

---
