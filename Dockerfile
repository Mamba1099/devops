# Multi-stage build for Go
FROM golang:1.23-alpine AS builder

# Install security updates
RUN apk update && apk upgrade --no-cache && \
    apk add --no-cache ca-certificates tzdata

WORKDIR /app

# Copy go mod files from app directory
COPY app/go.mod app/go.sum ./
RUN go mod download && go mod verify

# Copy source code from app directory
COPY app/ .

# Build with security flags
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags="-w -s -extldflags '-static'" \
    -a -installsuffix cgo \
    -o main ./cmd/

# Production stage
FROM gcr.io/distroless/static-debian12:latest

WORKDIR /app

# Copy CA certificates and binary
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/main .

EXPOSE 3000

HEALTHCHECK --interval=30s --timeout=3s --start-period=10s --retries=3 \
  CMD ["/app/main", "-health"]

CMD ["./main"]