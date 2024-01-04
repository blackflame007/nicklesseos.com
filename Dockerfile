# Build stage for Node.js
FROM node:21-slim as builder-node

# Set the working directory for Node.js
WORKDIR /app

# Install make
RUN apt-get update && apt-get install -y make

# Copy make file
COPY Makefile ./

# Copy the Node.js application and dependencies
COPY app/package.json app/package-lock.json ./

# Copy the frontend source code
COPY ./ ./

# Build the Node.js frontend
RUN make build-js

# Final stage for Go
FROM golang:1.21 as builder-go

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the Node.js frontend assets from the "builder-node" stage to the Go app directory
COPY --from=builder-node /app/app/assets/dist ./app/assets/dist

# Copy the rest of the source code
COPY ./ ./

# Generate HTML using templ and build the Go app
RUN make build

# Optional: Define the health check for the container
HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 CMD [ "curl", "-f", "http://localhost:3000/health" ]

# Document that the service listens on port 3000
EXPOSE 3000

# Run the binary using your Makefile's "run" target
CMD ["make", "run"]
