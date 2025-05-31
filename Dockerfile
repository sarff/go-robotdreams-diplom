FROM golang:1.24-alpine AS build

WORKDIR /work

COPY go.mod go.sum ./
RUN --mount=type=cache,id=gomod,target=/go/pkg/mod \
    --mount=type=cache,id=gobuild,target=/root/.cache/go-build \
    go mod download

COPY cmd/ cmd/
COPY internal internal

RUN --mount=type=cache,id=gobuild,target=/root/.cache/go-build \
    go build -o /tmp/bin/server ./cmd/server

FROM scratch AS deploy
COPY --from=build /tmp/bin /usr/local/bin
CMD ["/usr/local/bin/server"]