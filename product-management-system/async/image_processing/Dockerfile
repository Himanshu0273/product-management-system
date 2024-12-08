# Use the official Golang image as the base
FROM golang:1.19-alpine

# Set the Current Working Directory inside the container for the image processing service
WORKDIR /app

# Copy go.mod and go.sum to the working directory
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the entire content of the async/image_processing directory into the container
COPY . .

# Build the Go application (image processing service)
RUN go build -o processor ./main.go

# Expose the service port (if necessary, e.g., for communication or testing)
EXPOSE 8081

# Run the image processing service when the container starts
CMD ["./processor"]
