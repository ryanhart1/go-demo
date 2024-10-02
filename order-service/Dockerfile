FROM golang:1.23.1

WORKDIR /app/order-service

COPY . .
RUN go mod tidy

RUN go build -o order-service .
# RUN go run server.go

# Command to run the executable
# CMD ["./order-service"]

CMD ["go", "run", "server.go"]