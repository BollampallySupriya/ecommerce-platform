# Use the official Go image as the base image
FROM golang:1.22.7 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files first for better caching
COPY go.mod go.sum ./

# Download the necessary dependencies
RUN go mod download

# Copy the entire project files into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server/main.go

# Use a smaller image to run the app
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Copy the .env file to the root directory of the container
COPY --from=builder /app/.env .

# Set the permissions for the built binary
RUN chmod +x main

# Command to run the executable
CMD ["./main"]
