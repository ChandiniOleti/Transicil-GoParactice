# ==========================
# Stage 1 - Build
# ==========================
FROM golang:1.25 AS builder

WORKDIR /workspace

# Copy workspace
COPY go.work ./

# Copy all modules
COPY shared ./shared
COPY APIGateway ./APIGateway
COPY AuthService ./AuthService
COPY UserService ./UserService
COPY MerchantService ./MerchantService
COPY TransactionService ./TransactionService
COPY PaybackService ./PaybackService
COPY ReportService ./ReportService
COPY PayLaterService ./PayLaterService

# Build PaybackService
WORKDIR /workspace/PaybackService

RUN go mod download
RUN go build -o paybackservice .

# ==========================
# Stage 2 - Runtime
# ==========================
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /workspace/PaybackService/paybackservice .
COPY --from=builder /workspace/PaybackService/schema.sql .
COPY --from=builder /workspace/PaybackService/.env ./

EXPOSE 8085

CMD ["./paybackservice"]