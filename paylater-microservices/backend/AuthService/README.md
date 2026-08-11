# AuthService

Port: **8081**

Issues JWTs for user, admin, and merchant login. Hosts admin CRUD.

## Run

```bash
cp .env.example .env   # if needed
go run .
```

## Endpoints

- `POST /login`, `/admin/login`, `/merchant/login` (rate limited)
- `POST/GET /admins` (JWT + ADMIN)
- `GET /health`

## Env

`SERVICE_PORT`, `DB_*`, `JWT_SECRET`, `INTERNAL_API_TOKEN`
