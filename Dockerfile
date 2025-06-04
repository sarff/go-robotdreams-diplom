FROM golang:latest AS builder

WORKDIR /app

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build application
RUN CGO_ENABLED=0 GOOS=linux go build -o shatapp ./cmd/

FROM alpine:latest

# Install runtime dependencies
RUN apk --no-cache add ca-certificates curl

WORKDIR /root/

# Copy binary
COPY --from=builder /app/shatapp ./shatapp

# Set permissions
RUN chmod +x ./shatapp

EXPOSE 8081

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD curl -f http://localhost:8081/health || exit 1

ENTRYPOINT ["./shatapp"]