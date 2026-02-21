# ---------- Stage 1: Build ----------
FROM golang:1.25.5-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build static binary
RUN CGO_ENABLED=0 GOOS=linux go build -p 2 -ldflags="-s -w" -o app .

# ---------- Stage 2: Run ----------
FROM alpine:3.19

WORKDIR /app

# Add non-root user (security)
RUN adduser -D appuser

COPY --from=builder /app/app .

USER appuser

EXPOSE 8080

CMD ["./app"]
