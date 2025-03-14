# Use a modern Go version
FROM golang:1.24-alpine AS build

# Install dependencies
RUN apk --no-cache add gcc g++ make ca-certificates

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first
COPY go.mod go.sum ./

# Download dependencies before copying source files
RUN go mod tidy && go mod vendor

# Copy the entire project into the container
COPY . .

# Build the application
RUN go build -mod=vendor -o /go/bin/app ./account/cmd/account

# Create a minimal final image
FROM alpine:3.11
WORKDIR /usr/bin

# Copy built binary from the build stage
COPY --from=build /go/bin/app .

# Expose application port
EXPOSE 8000

# Start the application
CMD ["./app"]
