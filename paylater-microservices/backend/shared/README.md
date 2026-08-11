# shared

Reusable infrastructure only — no business logic.

```
shared/
├── config/      # dotenv + env helpers + DB DSN + secret masking
├── constants/   # header/role constants
├── error/       # common error message constants
├── httpclient/  # HTTP client factory
├── logger/      # structured slog + Gin access log
├── ratelimit/   # per-IP Gin rate limiter
├── requestid/   # X-Request-ID middleware
├── response/    # JSON helpers (preserves {"error":"..."} shapes)
└── server/      # graceful shutdown + /health handler
```

Module path: `paylater.dev/shared`
