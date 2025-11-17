# Step 1: Build stage
FROM golang:1.25.4-alpine3.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./v1/cmd

FROM alpine:3.22.1

WORKDIR /app
RUN apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Bangkok /etc/localtime \
    && echo "Asia/Bangkok" > /etc/timezone

COPY --from=builder /app/main .
RUN chmod +x main

CMD ["./main"]