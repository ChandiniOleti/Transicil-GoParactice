# APIGateway

Port: **8080**

Single client entry point. Validates JWT on protected routes and reverse-proxies to downstream services. Does not issue tokens. Does not rate-limit forwarding (public limits live on Auth/User/Merchant services).

## Run

```bash
go run .
```

## Env

`SERVICE_PORT`, `JWT_SECRET`, `AUTH_SERVICE_URL`, `USER_SERVICE_URL`, `MERCHANT_SERVICE_URL`, `TRANSACTION_SERVICE_URL`, `PAYBACK_SERVICE_URL`, `REPORT_SERVICE_URL`
