FROM golang:1.21-alpine as builder

WORKDIR /app

COPY ./cmd/email-scheduler/ .
COPY ./internal/ ./internal
COPY go.mod go.sum .env ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -o main .

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder app/main main
COPY --from=builder app/.env .env

ENTRYPOINT ["./main"]