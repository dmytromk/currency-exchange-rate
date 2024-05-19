FROM golang:1.21-alpine as builder

WORKDIR /app

COPY ./cmd/api-server/ .
COPY ./internal/ ./internal
COPY go.mod go.sum ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -o main .

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder app/main main

ENTRYPOINT ["./main"]