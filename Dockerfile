# ---- Build stage ----
FROM golang:1.25.5-alpine AS builder

WORKDIR /app

# Faster + reproducible builds
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build the CLI binary (cmd/tmdb)
RUN CGO_ENABLED=0 GOOS=linux go build -o /tmdb ./cmd/tmdb

# ---- Runtime stage ----
FROM alpine:3.20

# (Optional) root CA certs for HTTPS
RUN apk add --no-cache ca-certificates

COPY --from=builder /tmdb /usr/local/bin/tmdb

ENTRYPOINT ["tmdb"]
