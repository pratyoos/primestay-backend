# PrimeStay Backend

**Enterprise Rental & Housing Management System**

PrimeStay is an enterprise-grade rental and housing management platform designed to manage multiple housing organizations, residential projects, buildings, apartments, tenants, leases, billing, payments, and property operations.

The backend is developed using **Go (Golang)** with a **Modular Monolith and Event-Driven Architecture**, providing RESTful APIs for web applications, mobile applications, and third-party integrations.

## Architecture

- **Architecture:** Modular Monolith + Event-Driven Architecture
- **Language:** Go (Golang)
- **HTTP Framework:** Chi
- **API:** REST
- **Documentation:** OpenAPI / Swagger
- **Database:** PostgreSQL
- **Database Driver:** pgx
- **SQL Generation:** sqlc
- **Migrations:** Goose
- **Caching:** Redis
- **Event Streaming:** Apache Kafka (KRaft)
- **Kafka Client:** franz-go
- **Logging:** slog
- **Containerization:** Docker / Docker Compose
- **Frontend:** Separate repository

## Technology Responsibilities

| Technology | Purpose |
|---|---|
| Go | Backend application |
| Chi | HTTP routing |
| PostgreSQL | Persistent business data |
| pgx | Database connectivity |
| sqlc | Type-safe SQL generation |
| Goose | Database migrations |
| Redis | Caching, sessions and temporary data |
| Kafka | Event streaming and asynchronous processing |
| franz-go | Kafka producers and consumers |
| OpenAPI | API documentation |
| slog | Structured logging |
| Docker | Containerization |

## Project Structure

```text
primestay-backend/
├── cmd/
│   ├── api/
│   │   └── main.go
│   └── worker/
│       └── main.go
├── internal/
│   ├── app/
│   ├── platform/
│   │   ├── config/
│   │   ├── database/
│   │   ├── redis/
│   │   ├── kafka/
│   │   ├── logger/
│   │   └── middleware/
│   ├── modules/
│   │   ├── auth/
│   │   ├── organization/
│   │   ├── property/
│   │   ├── building/
│   │   ├── unit/
│   │   ├── tenant/
│   │   ├── lease/
│   │   ├── billing/
│   │   ├── payment/
│   │   ├── maintenance/
│   │   └── report/
│   ├── workers/
│   │   ├── notification/
│   │   ├── analytics/
│   │   └── event/
│   └── shared/
├── api/
│   └── openapi.yaml
├── migrations/
├── queries/
├── pkg/
├── scripts/
├── tests/
├── deployments/
├── docs/
├── .env.example
├── .gitignore
├── Dockerfile
├── docker-compose.yml
├── Makefile
├── sqlc.yaml
├── go.mod
├── go.sum
└── README.md
```

## Core Modules

| Module | Responsibility |
|---|---|
| Authentication | Login, registration, sessions and access control |
| Organization | Housing organization management |
| Property | Residential project management |
| Building | Buildings, blocks and floors |
| Unit | Apartments, rooms and commercial units |
| Tenant | Tenant onboarding and profiles |
| Lease | Rental agreements, renewals and termination |
| Billing | Rent generation, invoices and balances |
| Payment | Payment processing and verification |
| Maintenance | Maintenance requests and operations |
| Reports | Operational and financial reporting |

## Kafka Event Architecture

Kafka enables asynchronous communication between application modules, workers, and external integrations.

**Initial Topics:**

| Topic | Events |
|---|---|
| tenant.events | tenant.created, tenant.updated |
| lease.events | lease.created, lease.expiring |
| billing.events | invoice.generated, invoice.overdue |
| payment.events | payment.completed, payment.failed |
| maintenance.events | maintenance.created, maintenance.resolved |
| notification.events | notification.requested |

**Event Infrastructure:**
- Common event envelope and schema versioning
- Kafka producers and consumers
- Consumer groups
- Transactional Outbox Pattern
- Idempotent event processing
- Retry and dead-letter handling

Critical business events, particularly payments and billing, will use the Transactional Outbox Pattern to ensure reliable event publishing.

## Architectural Principles

1. Modular Monolith with clear module boundaries.
2. Event-driven asynchronous processing using Kafka.
3. PostgreSQL as the primary source of truth.
4. Redis for caching and temporary data.
5. Separation of handlers, business logic and repositories.
6. API-first development using OpenAPI.
7. Multi-tenancy with organization-level data isolation.
8. Secure authentication, authorization and auditability.
9. Idempotent financial operations and event processing.
10. Scalable background workers and integrations.

## Development Setup

**Prerequisites:** Go, Docker Desktop, Git.

### 1. Clone Repository

```bash
git clone https://github.com/YOUR_USERNAME/primestay-backend.git
cd primestay-backend
```

### 2. Configure Environment

```bash
cp .env.example .env
```

### 3. Start Infrastructure

```bash
docker compose up -d
```

Services: PostgreSQL, Redis and Kafka (KRaft).

### 4. Install Dependencies

```bash
go mod tidy
```

### 5. Run Migrations

```bash
goose -dir migrations postgres "$DATABASE_URL" up
```

### 6. Generate SQL Code

```bash
sqlc generate
```

### 7. Start Application

```bash
go run ./cmd/api
```

Start background workers:

```bash
go run ./cmd/worker
```

## Database & API

- **Migrations:** Goose
- **Database Queries:** sqlc
- **Database:** PostgreSQL
- **API Specification:** `api/openapi.yaml`
- **API Versioning:** `/api/v1/`

## Testing

```bash
go test ./...
go test -v ./...
go test -race ./...
go test -cover ./...
```

## Repository Strategy

| Repository | Responsibility |
|---|---|
| primestay-backend | Go backend, APIs, business logic, databases and event processing |
| primestay-frontend | Web application and user interface |

## Development Roadmap

- [ ] Go project initialization and architecture
- [ ] Application configuration, logging and Chi router
- [ ] PostgreSQL, pgx and Goose integration
- [ ] sqlc configuration
- [ ] Redis integration
- [ ] Kafka and franz-go integration
- [ ] Event envelope and topic configuration
- [ ] Transactional Outbox and event consumers
- [ ] Authentication and organization management
- [ ] Property, building and unit management
- [ ] Tenant and lease management
- [ ] Billing and payment management
- [ ] Notification and analytics workers
- [ ] Maintenance and reporting
- [ ] API documentation and testing
- [ ] Security hardening and production deployment

## Project Status

**Status:** Initial Development

PrimeStay is currently in its foundation and architecture development stage. Infrastructure and core business modules will be implemented incrementally.

---