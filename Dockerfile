FROM alpine:latest

# Add CA certificates
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /app

# Copy the pre-built binary
COPY bin/main .

# Run the binary
ENTRYPOINT ["./main"]