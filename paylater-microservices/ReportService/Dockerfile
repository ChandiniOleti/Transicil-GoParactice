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

# Build ReportService
WORKDIR /workspace/ReportService

RUN go mod download
RUN go build -o reportservice .

# ==========================
# Stage 2 - Runtime
# ==========================
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /workspace/ReportService/reportservice .
COPY --from=builder /workspace/ReportService/.env ./

EXPOSE 8086

CMD ["./reportservice"]