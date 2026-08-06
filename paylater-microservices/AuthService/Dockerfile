# ==========================
# Stage 1 - Build
# ==========================
FROM golang:1.25 AS builder

WORKDIR /workspace

# Copy workspace file
COPY go.work ./

# Copy all workspace modules
COPY shared ./shared
COPY APIGateway ./APIGateway
COPY AuthService ./AuthService
COPY UserService ./UserService
COPY MerchantService ./MerchantService
COPY TransactionService ./TransactionService
COPY PaybackService ./PaybackService
COPY ReportService ./ReportService
COPY PayLaterService ./PayLaterService

# Build AuthService
WORKDIR /workspace/AuthService

RUN go mod download
RUN go build -o authservice .


# ==========================
# Stage 2 - Runtime
# ==========================
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /workspace/AuthService/authservice .

COPY --from=builder /workspace/AuthService/schema.sql .
COPY --from=builder /workspace/AuthService/.env ./

EXPOSE 8081

CMD ["./authservice"]