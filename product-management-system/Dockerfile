# Use official Golang image as the base for the API service
FROM golang:1.19-alpine

# Set the Current Working Directory inside the container (API service)
WORKDIR /app

# Copy go.mod and go.sum to the working directory for dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the entire project into the container (this assumes 'api/' is at the root)
COPY . .

# Build the Go application (main entry point for the API)
RUN go build -o main ./api

# Expose the API service on port 8080
EXPOSE 8080

# Run the Go executable when the container starts (start the API service)
CMD ["./main"]
