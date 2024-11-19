# Use the official Go image as the base image
FROM golang:1.22 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files (go.mod and go.sum)
COPY go.mod go.sum ./

# Download and cache the dependencies
RUN go mod download

# Copy the workspace code
COPY . .

# Build the Go application
RUN go build -o chronicle .

# Use a smaller base image for the final image
FROM scratch

# Copy the built binary to the final image
COPY --from=builder /app/chronicle /chronicle

# Set the entrypoint to execute the binary when the container starts
ENTRYPOINT ["/chronicle"]