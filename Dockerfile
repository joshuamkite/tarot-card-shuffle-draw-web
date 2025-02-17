# Start from a base image with Go installed
FROM golang:1.24 AS builder

# Set the working directory
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Tidy Go modules
RUN go mod tidy

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/tarot_shuffle_draw main.go

# Create user and group files that will be copied to the final 'from scratch' image
RUN mkdir /user && \
    echo 'tarot-app:x:10000:10000::/app:' > /user/passwd && \
    echo 'tarot-app:x:10000:' > /user/group

# Start a new stage from scratch
FROM scratch

# Copy the user and group files
COPY --from=builder /user/passwd /etc/passwd
COPY --from=builder /user/group /etc/group

# Set environment variable to suppress debug messages
ENV GIN_MODE=release

# Set default port (can be overridden at runtime)
ENV PORT=8080

# Copy the binary from the builder stage
COPY --from=builder /app/tarot_shuffle_draw /tarot_shuffle_draw

# Copy static files
COPY --from=builder /app/static /static
COPY --from=builder /app/templates /templates

# Switch to non-root user
USER tarot-app

# Expose the port from the environment variable
EXPOSE 8080

# Add a label for the image description
LABEL org.opencontainers.image.description="Tarot card shuffle draw web application"

# Command to run the executable
CMD ["/tarot_shuffle_draw"]