# AuthService

Port: **8081**

Issues JWTs for user, admin, and merchant login. Hosts admin CRUD.

## Run

```bash
cp .env.example .env   # if needed
go run .
```

## Endpoints

- `GET /login/public-key` — RSA public key for encrypted login payloads
- `POST /login`, `/admin/login`, `/merchant/login` (rate limited; encrypted credentials only)
- `POST/GET /admins` (JWT + ADMIN)
- `GET /health`

## Env

`SERVICE_PORT`, `DB_*`, `JWT_SECRET`, `INTERNAL_API_TOKEN`, `LOGIN_RSA_PRIVATE_KEY`

Generate a local login key with:

```bash
cd AuthService
go run scripts/setup_login_key.go
```
