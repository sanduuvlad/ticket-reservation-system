FROM golang:1.26.1-alpine

WORKDIR /app

COPY . .

RUN go build -o event-service ./cmd/event-service

CMD ["./event-service"]