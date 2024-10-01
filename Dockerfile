FROM docker.io/golang:1.22-bookworm AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN make build

FROM docker.io/debian:12-slim AS server

COPY --from=builder /app/microblog-api /app/microblog-api

EXPOSE 1323

CMD ["/app/microblog-api"]
