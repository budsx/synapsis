FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -installsuffix cgo -o order-service
RUN ls -all && pwd

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /app/order-service /order-service

ENTRYPOINT ["/order-service"]