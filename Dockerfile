# Build stage
FROM golang:1.21 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Generate HTML using templ
RUN go run github.com/a-h/templ/cmd/templ@latest generate

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/nicklesseos.com/main.go



# Final stage: Use a lightweight base image
FROM alpine:latest

# Create a user and group for running the application
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Set the working directory in the container
WORKDIR /home/appuser

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Use the non-root user
USER appuser

# Optional: Define the health check for the container
HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 CMD [ "curl", "-f", "http://localhost:3000/health" ]

# Document that the service listens on port 3000
EXPOSE 3000

# Run the binary
CMD ["./main"]
