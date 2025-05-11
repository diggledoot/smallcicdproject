# Build
FROM golang:1.24.3-bookworm AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o app_binary

# Run
FROM alpine:3.21.3
COPY --from=builder /app/app_binary /
EXPOSE 8080
CMD ["/app_binary"]
