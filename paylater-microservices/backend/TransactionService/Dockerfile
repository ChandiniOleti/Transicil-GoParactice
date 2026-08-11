# ==========================
# Stage 1 - Build
# ==========================
FROM golang:1.25 AS builder

WORKDIR /workspace

COPY go.work ./

COPY shared ./shared
COPY APIGateway ./APIGateway
COPY AuthService ./AuthService
COPY UserService ./UserService
COPY MerchantService ./MerchantService
COPY TransactionService ./TransactionService
COPY PaybackService ./PaybackService
COPY ReportService ./ReportService
COPY PayLaterService ./PayLaterService

WORKDIR /workspace/TransactionService

RUN go mod download
RUN go build -o transactionservice .

# ==========================
# Stage 2 - Runtime
# ==========================
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /workspace/TransactionService/transactionservice .
COPY --from=builder /workspace/TransactionService/schema.sql .
COPY --from=builder /workspace/TransactionService/.env ./

EXPOSE 8084

CMD ["./transactionservice"]