# UserService

Port: **8082**

User CRUD and internal due updates for Transaction/Payback/Report.

## Run

```bash
go run .
```

## Endpoints

- `POST /users` (public, rate limited)
- `GET/PUT/DELETE /users/:id`, `GET /users` (JWT)
- Internal: `/internal/users`, `/internal/users/:id`, `PATCH /internal/users/:id/due`
- `GET /health`

## Env

`SERVICE_PORT`, `DB_*`, `JWT_SECRET`, `INTERNAL_API_TOKEN`
