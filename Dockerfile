# Step 1: Build the Go application
FROM golang:1.23 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the container and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go application for production
RUN go build -o main .

# Step 2: Run the Go application in a lightweight image
FROM debian:stable-slim

## Set environment variables for Heroku
#ENV PORT=8080
#ENV GO_ENV=production

# Set the working directory inside the container
WORKDIR /app

# Copy the built executable from the builder stage
COPY --from=builder /app/main .

COPY /pkg/config /app/pkg/config

# Expose the port that the application will listen on
EXPOSE 8080

# Command to run the Go application
CMD ["./main.go"]
