# TransactionService

Port: **8084**

Purchase orchestration and transaction history. Calls User and Merchant over REST.

## Run

```bash
go run .
```

## Endpoints

- `POST /transactions`, history routes (JWT)
- Internal payback create + list APIs
- `GET /health`

## Env

`SERVICE_PORT`, `DB_*`, `JWT_SECRET`, `INTERNAL_API_TOKEN`, `USER_SERVICE_URL`, `MERCHANT_SERVICE_URL`
