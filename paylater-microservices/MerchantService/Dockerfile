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

WORKDIR /workspace/MerchantService

RUN go mod download
RUN go build -o merchantservice .

# ==========================
# Stage 2 - Runtime
# ==========================
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /workspace/MerchantService/merchantservice .
COPY --from=builder /workspace/MerchantService/schema.sql .
COPY --from=builder /workspace/MerchantService/.env ./

EXPOSE 8083

CMD ["./merchantservice"]