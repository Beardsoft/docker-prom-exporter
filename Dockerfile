# Use the official Go image to create a build artifact.
FROM golang:1.19 as builder

# Create and change to the app directory.
WORKDIR /app

# Copy the Go Modules manifests and download the dependencies
COPY src/go.mod src/go.sum ./
RUN go mod download

# Copy the source code into the container.
COPY src/ ./

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o docker-monitor

# Use the Alpine image for the runtime container.
FROM alpine:latest
WORKDIR /root/

# Copy the compiled binary from the builder stage.
COPY --from=builder /app/docker-monitor .

# Command to run the binary.
CMD ["./docker-monitor"]
