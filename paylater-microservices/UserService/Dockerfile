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

WORKDIR /workspace/UserService

RUN go mod download
RUN go build -o userservice .

# ==========================
# Stage 2 - Runtime
# ==========================
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /workspace/UserService/userservice .
COPY --from=builder /workspace/UserService/schema.sql .
COPY --from=builder /workspace/UserService/.env ./

EXPOSE 8082

CMD ["./userservice"]