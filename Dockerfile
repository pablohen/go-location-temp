FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go.mod file
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server ./cmd/api

# Use a smaller image for the final container
FROM alpine:latest

# Add CA certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/server /app/server

# Expose the port
EXPOSE 8080

# Set the entry point
CMD ["/app/server"]
