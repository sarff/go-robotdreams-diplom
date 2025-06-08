FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN --mount=type=cache,id=gomod,target=/go/pkg/mod \
    --mount=type=cache,id=gobuild,target=/root/.cache/go-build \
    go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o shatapp ./cmd/

FROM alpine:latest

RUN apk --no-cache add ca-certificates curl

WORKDIR /app

COPY --from=builder /app/shatapp ./shatapp
COPY --from=builder /app/static ./static
RUN chmod +x ./shatapp

EXPOSE 8081

ENTRYPOINT ["./shatapp"]