# Start from a base image with Go installed
FROM golang:1.22 as builder

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

# Start a new stage from scratch
FROM scratch

# Set environment variable to suppress debug messages
ENV GIN_MODE=release

# Copy the binary from the builder stage
COPY --from=builder /app/tarot_shuffle_draw /tarot_shuffle_draw

# Copy static files
COPY --from=builder /app/static /static
COPY --from=builder /app/templates /templates

# Expose the port the app runs on
EXPOSE 80

# Add a label for the image description
LABEL org.opencontainers.image.description="Tarot card shuffle draw web application"

# Command to run the executable
CMD ["/tarot_shuffle_draw"]
