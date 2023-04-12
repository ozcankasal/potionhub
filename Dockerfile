# Use the official Golang image as a base image
FROM golang:1.17-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the workspace
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the entire source code from the local context to the working directory
COPY . .

# Build the application
RUN go build -o main .

# Expose the port the app will run on
EXPOSE 8080

# Command to run the application when the container starts
CMD ["/app/main"]
