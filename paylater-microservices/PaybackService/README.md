# PaybackService

Port: **8085**

Payback orchestrator (no local SQLC). Uses User + Transaction internal APIs.

## Run

```bash
go run .
```

## Endpoints

- `POST /payback` (JWT)
- `GET /health`

## Env

`SERVICE_PORT`, `JWT_SECRET`, `INTERNAL_API_TOKEN`, `USER_SERVICE_URL`, `TRANSACTION_SERVICE_URL`
