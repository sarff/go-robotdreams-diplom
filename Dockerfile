FROM golang:1.24 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN --mount=type=cache,id=gomod,target=/go/pkg/mod \
    --mount=type=cache,id=gobuild,target=/root/.cache/go-build \
    go mod download
RUN go build -o chatapp ./cmd/server
FROM debian:bookworm-slim

WORKDIR /app
COPY --from=builder /app/chatapp .

EXPOSE 8081

ENTRYPOINT ["./chatapp"]