# PayLater — Microservices Platform

Production-ready BNPL (Buy Now, Pay Later) platform built with Go microservices.

## Architecture

```
                     Clients
                        |
                        v
               +----------------+
               |  API Gateway   |
               |     :8080      |
               +--------+-------+
                        |
     +---------+--------+--------+---------+---------+
     |         |        |        |         |         |
     v         v        v        v         v         v
  Auth:8081  User:8082 Merchant Transaction Payback  Report
                       :8083      :8084      :8085    :8086
     |         |        |          |                   
     +----+----+--------+----------+----> MySQL (shared)
```

## Tech Stack

- Go + Gin
- SQLC + MySQL
- JWT + bcrypt
- REST inter-service communication
- Shared infra module (`paylater.dev/shared`)

## Ports

| Service | Port |
|---------|------|
| APIGateway | 8080 |
| AuthService | 8081 |
| UserService | 8082 |
| MerchantService | 8083 |
| TransactionService | 8084 |
| PaybackService | 8085 |
| ReportService | 8086 |

## Folder Structure

```
PayLater/
├── APIGateway/
├── AuthService/
├── UserService/
├── MerchantService/
├── TransactionService/
├── PaybackService/
├── ReportService/
├── shared/                 # reusable infra only
│   ├── config/
│   ├── constants/
│   ├── error/
│   ├── httpclient/
│   ├── logger/
│   ├── ratelimit/
│   ├── requestid/
│   ├── response/
│   └── server/
├── go.work
├── .gitignore
└── README.md
```

## Service Responsibilities

| Service | Responsibility |
|---------|----------------|
| APIGateway | Single entry point; JWT gate for protected routes; reverse proxy |
| AuthService | Login (user/admin/merchant), admin CRUD, JWT issuance |
| UserService | User CRUD + internal due updates |
| MerchantService | Merchant CRUD + commission + internal commission lookup |
| TransactionService | Purchases + history; orchestrates User/Merchant via REST |
| PaybackService | Payback orchestration via User + Transaction internals |
| ReportService | Aggregated reports via User + Transaction internals |

## Authentication Flow

1. Client calls `POST /login` (or admin/merchant login) via Gateway `:8080`
2. AuthService validates credentials and returns JWT
3. Client sends `Authorization: Bearer <token>` on protected routes
4. Gateway validates JWT, then forwards to downstream service
5. Downstream services may re-validate JWT for defense in depth
6. Service-to-service calls use `X-Internal-Token`

## REST Communication

- Transaction → User (JWT) + Merchant/User (internal)
- Payback → User (JWT/internal) + Transaction (internal)
- Report → User + Transaction (internal)

## How to Run


## Run with Docker

### Prerequisites

- Docker
- Docker Compose

### Clone the repository

```bash
git clone https://github.com/<your-username>/<your-repository>.git
cd paylater-microservices
```

### Configure Environment Files

Copy each `.env.example` file as `.env`.

Linux / WSL:

```bash
cp AuthService/.env.example AuthService/.env
cp UserService/.env.example UserService/.env
cp MerchantService/.env.example MerchantService/.env
cp TransactionService/.env.example TransactionService/.env
cp PaybackService/.env.example PaybackService/.env
cp ReportService/.env.example ReportService/.env
cp APIGateway/.env.example APIGateway/.env
```

Update the values inside each `.env` file.

### Start the application

```bash
docker compose up -d
```

### Check running containers

```bash
docker compose ps
```

### Stop all containers

```bash
docker compose down
```


## Docker Hub Images

The following pre-built Docker images are available:

- chandini9/apigateway
- chandini9/authservice
- chandini9/userservice
- chandini9/merchantservice
- chandini9/transactionservice
- chandini9/paybackservice
- chandini9/reportservice

### Prerequisites

- Go 1.25+
- MySQL with database `paylaterdb`
- Copy each service's `.env.example` → `.env` and set secrets

### Start services (separate terminals)

```bash
cd AuthService && go run .
cd UserService && go run .
cd MerchantService && go run .
cd TransactionService && go run .
cd PaybackService && go run .
cd ReportService && go run .
cd APIGateway && go run .
```

Clients should call **only** `http://localhost:8080`.

### Health checks

```bash
curl http://localhost:8080/health
curl http://localhost:8081/health
# ... each service /health
```

Expected:

```json
{"status":"UP","service":"APIGateway"}
```

## Environment Variables

Common:

- `SERVICE_NAME`, `SERVICE_PORT`
- `JWT_SECRET`
- `INTERNAL_API_TOKEN` (services with internal APIs / callers)
- `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME` (DB-backed services)

Gateway / orchestrators also need:

- `AUTH_SERVICE_URL`, `USER_SERVICE_URL`, `MERCHANT_SERVICE_URL`
- `TRANSACTION_SERVICE_URL`, `PAYBACK_SERVICE_URL`, `REPORT_SERVICE_URL`

## Production Features

- Env-based configuration (fail-fast on missing secrets)
- Structured JSON logging with request IDs
- Graceful shutdown (SIGINT/SIGTERM)
- Health endpoints with DB ping where applicable
- Rate limiting on public login/registration endpoints
- Consistent error helpers in `shared/response`
- Go workspace (`go.work`) for multi-module development

## Build

From repo root (workspace enabled):

```bash
cd AuthService && go build -o bin/AuthService.exe .
# repeat per service
```
