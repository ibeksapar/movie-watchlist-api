FROM golang:1.23-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./

COPY user-service/go.mod user-service/go.sum ./user-service/

RUN go mod download

COPY . .

WORKDIR /app/user-service
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/user-service-main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/user-service-main ./main
EXPOSE 8081
ENTRYPOINT ["./main"]
