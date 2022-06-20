FROM golang:1.18 as builder

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . ./

RUN go build -o server cmd/main.go && \
    go build -o cli cmd/cli/main.go

FROM debian:buster-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/server /app/server

CMD ["/app/server"]