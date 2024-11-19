# Use the official Go image as the base image
FROM golang:1.22 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy chronicle code, include go.mod and go.sum
COPY . .

# Download and cache the dependencies
# TODO: Manually adding internal common here, need to figure out the best way
# to automatically install it via workspaces
RUN go mod download
RUN go get github.com/SomethingSexy/chronicle/internal/common@latest

# Build the Go application
RUN go build -o chronicle

# Use a smaller base image for the final image
# bookworm includes GLIBC
FROM debian:bookworm-slim

# Set up a non-root user (optional, but recommended for security)
RUN useradd -m appuser

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/chronicle .

# Change ownership of the binary (if using a non-root user)
RUN chown appuser:appuser /app/chronicle

# Switch to the non-root user
USER appuser

# Command to run the binary
CMD ["./chronicle"]
