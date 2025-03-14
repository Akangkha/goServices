# Use a modern Go version
FROM golang:1.24-alpine AS build

# Install necessary packages
RUN apk --no-cache add gcc g++ make ca-certificates

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum first for better caching
COPY go.mod go.sum ./

# Download dependencies before copying source files
RUN go mod tidy && go mod vendor

# Copy the entire project (all source files)
COPY . .

# Build the 'orders' service binary
RUN go build -mod=vendor -o /go/bin/app ./orders/cmd/order

# Use a minimal final image
FROM alpine:3.11
WORKDIR /usr/bin

# Copy the built binary from the build stage
COPY --from=build /go/bin/app .

# Expose the application port
EXPOSE 8000

# Start the application
CMD ["./app"]
