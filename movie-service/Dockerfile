FROM golang:1.23-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./

COPY movie-service/go.mod movie-service/go.sum ./movie-service/

RUN go mod download

COPY . .

WORKDIR /app/movie-service
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/movie-service-main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/movie-service-main ./main
EXPOSE 8080
ENTRYPOINT ["./main"]
