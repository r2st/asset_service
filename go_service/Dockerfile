# Builder stage
FROM golang:1.22.3-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main .

# Final stage
FROM alpine:latest
WORKDIR /root/
# Copy the built executable and the config file from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/config.json .
CMD ["./main"]
