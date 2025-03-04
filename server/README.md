# Archium - Server

In order to understand how the Archium platform works behind the scenes, we need to define some key concepts first.

### Diagram

A Diagram represents a complete system architecture simulation. A Diagram is composed of interconnected **Workers** that define service interactions under the oversight of the **Orchestrator**. A single Diagram can model a cloud infrastructure, a microservices ecosystem, or any distributed system, allowing users to analyze its performance, behavior, and cost estimations.

Running a Diagram triggers a real-time simulation, processing defined constraints like latency, failure rates, and request load.

![Architecture](../docs/diagram.png)

### Orchestrator

The Orchestrator is responsible for executing **Diagrams**. It manages the lifecycle of services, gathers simulation data, and reports results to the frontend.

- Spawns and manages Workers.
- Collects and aggregates simulation data.
- Sends structured reports and metrics to the frontend.

### Worker

A worker represents an individual service within a **Diagram**. Each worker has custom properties:

- Latency (expected delay in processing requests).
- Failure rate (probability of errors occurring).
- Request handling capacity (simulated throughput).
- Database size (if applicable, estimated storage and query operations).

## Analysis

Each service instance is independent, with its own properties like latency, failure rate, and throughput. Therefore, each service definition is easy to modify, test, and extend without breaking everything.

The Orchestrator acts as a manager, collecting events from services and sending aggregated data to the frontend.

Services communicate with the Orchestrator via channels instead of directly sending updates to the frontend.

WebSockets (or another streaming method) are handled by the Orchestrator, ensuring controlled and batched updates to the frontend.

Each service runs as a separate Goroutine, making everything concurrent and highly scalable.

## Implementation plan

1. Each service runs as a Goroutine and sends metrics to a channel.
2. The Orchestrator listens to these channels and aggregates data.
3. Every X milliseconds, the Orchestrator sends a batch update to the frontend via WebSocket.

## Benefits

- **Concurrent**. Every service runs independently in Goroutines.

- **Decoupled**. Services don't know about each other, only the Orchestrator.

- **Efficient communication**. Using channels to avoid unnecessary locks.

- **Optimized**. Orchestrator batches and sends controlled updates.
