FROM golang:1.23.1

WORKDIR /app/product-service

COPY . .
RUN go mod tidy

RUN go build -o product-service .

# Command to run the executable
CMD ["./product-service"]