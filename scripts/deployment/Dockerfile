# Final image
FROM alpine:latest

# Add a lightweight runtime
RUN apk --no-cache add ca-certificates

# Set working directory in final image
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY / .

# Expose default app port
EXPOSE 8080

# Command to run the application
CMD ["./app"]
