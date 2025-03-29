# ───────────────────────────────────────────────────────────
# :construction_site: Stage 1: Build the Go application
# ───────────────────────────────────────────────────────────
FROM golang:1.23.3-alpine AS builder

# Install dependencies
RUN apk add --no-cache tzdata

# Set timezone to Asia/Bangkok
ENV TZ=Asia/Bangkok

# Set up working directory
WORKDIR /app

# Copy go.mod and go.sum first to cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code
COPY . .

# Build the Go application
RUN go build -o buffet-pos-backend .

# ───────────────────────────────────────────────────────────
# :rocket: Stage 2: Create the final runtime image
# ───────────────────────────────────────────────────────────
FROM alpine:latest

# Install necessary system packages
RUN apk --no-cache add ca-certificates

# Set up working directory
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/buffet-pos-backend .

# Ensure the binary is executable
RUN chmod +x buffet-pos-backend

# Copy timezone info for accurate timestamps
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /zoneinfo.zip
ENV ZONEINFO=/zoneinfo.zip

# Cloud Run requires the app to listen on $PORT, default to 8080
ENV PORT=8080

# Expose the port (for documentation, Cloud Run ignores it)

EXPOSE $PORT

# Start the Go application
CMD ["./buffet-pos-backend"]