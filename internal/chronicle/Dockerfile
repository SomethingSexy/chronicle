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
RUN go get github.com/SomethingSexy/chronicle/internal/common

# Build the Go application
RUN go build -o chronicle .

# Use a smaller base image for the final image
FROM scratch

# Copy the built binary to the final image
COPY --from=builder /app/chronicle /chronicle

# Set the entrypoint to execute the binary when the container starts
ENTRYPOINT ["/chronicle"]