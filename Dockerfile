# Base image
FROM golang:1.20 as builder

# Set working directory inside the container
WORKDIR /app

# Copy the application code
COPY . .

# Download dependencies and build the Go binary
RUN go mod tidy && go build -o app .

# Final image
FROM alpine:latest

# Add a lightweight runtime
RUN apk --no-cache add ca-certificates

# Set working directory in final image
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/app .

# Expose default app port
EXPOSE 8080

# Command to run the application
CMD ["./app"]
