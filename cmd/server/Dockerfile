# Use official Golang image as the base image
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY ../../go.mod ../../go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code from the current directory to the Working Directory inside the container
COPY . .

# Build the Go server
RUN go build -o server main.go

# Expose port 50051 to the outside world
EXPOSE 50051

# Command to run the server executable
CMD ["./server"]
