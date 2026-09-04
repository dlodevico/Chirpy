# Chirpy

Chirpy is a secure, high-performance RESTful API service built in Go. It features full user lifecycle management, JWT-based authentication with persistent refresh tokens, role-based account upgrades via external webhook integration, and database-level content filtering.

---

## Features

- **Authentication & Authorization**
  - Secure password hashing using **Argon2id**.
  - Short-lived JWT access tokens (`1-hour` expiration).
  - Opaque, cryptographically secure refresh tokens stored in PostgreSQL (`60-day` expiration) supporting session revocation.
- **Content Management**
  - Full CRUD lifecycle for Chirps with input sanitization and length bounds (`140 characters`).
  - Strict resource ownership checks (users can only delete their own content).
- **Filtering & Sorting**
  - Database-level filtering by `author_id` using optimized `sqlc` queries.
  - Flexible in-memory response sorting via query parameters (`sort=asc` | `sort=desc`).
- **External Webhooks**
  - Hardened `/api/polka/webhooks` endpoint for third-party subscription events (`is_chirpy_red`).
  - Custom HTTP header header API Key authentication (`Authorization: ApiKey <KEY>`).
- **Admin & Monitoring**
  - Traffic hit counters using lock-free atomic operations (`sync/atomic`).
  - Isolated environment resets (`PLATFORM=dev`).

---

## Tech Stack

- **Language:** Go 1.22+
- **Database:** PostgreSQL 15+
- **Database Access:** `sqlc` (type-safe SQL compiler)
- **Database Migrations:** `goose`
- **Configuration:** `godotenv`

---

## Getting Started

### Prerequisites

- [Go 1.22+](https://go.dev/dl/) installed.
- [PostgreSQL](https://www.postgresql.org/download/) running locally or via Docker.
- [goose](https://github.com/pressly/goose) for executing database migrations.
- [sqlc](https://sqlc.dev/) for generating Go code from raw SQL.

### Environment Setup

Create a `.env` file in the project root:

```env
PORT=8080
DB_URL=postgres://postgres:postgres@localhost:5432/chirpy?sslmode=disable
PLATFORM=dev
JWT_SECRET=your_jwt_secret_key_here
POLKA_KEY=f271c81ff7084ee5b99a5091b42d486e
