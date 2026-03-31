FROM golang:1.23-alpine AS builder

# Set working directory
WORKDIR /app

# Install git for go mod download
RUN apk add --no-cache git

# Copy go mod files first for better layer caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary with static linking
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static" -w -s' -o /api-template main.go


# Production stage
FROM alpine:latest

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Create non-root user
RUN addgroup -g 1001 -f app && adduser -S app -u 1001

# Create app directory
WORKDIR /root/

# Copy binary from builder stage
COPY --from=builder /api-template .

# Change ownership to non-root user
RUN chown -R app:app /root
USER app:app

# Expose port
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --quiet --tries=1 --spider http://localhost:8080/health || exit 1

# Start the binary
ENTRYPOINT ["./api-template"]