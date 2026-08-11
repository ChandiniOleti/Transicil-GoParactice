# MerchantService

Port: **8083**

Merchant CRUD, commission updates, and internal commission lookup.

## Run

```bash
go run .
```

## Endpoints

- `POST /merchants` (public, rate limited)
- Merchant/Admin JWT routes for CRUD + commission
- `GET /internal/merchants/:id/commission`
- `GET /health`

## Env

`SERVICE_PORT`, `DB_*`, `JWT_SECRET`, `INTERNAL_API_TOKEN`
