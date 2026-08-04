# ReportService

Port: **8086**

Admin report aggregator over User + Transaction internal APIs.

## Run

```bash
go run .
```

## Endpoints

- `/reports/*` (JWT + ADMIN)
- `GET /health`

## Env

`SERVICE_PORT`, `JWT_SECRET`, `INTERNAL_API_TOKEN`, `USER_SERVICE_URL`, `TRANSACTION_SERVICE_URL`
