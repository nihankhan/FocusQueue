# Build stage
FROM golang:1.20-alpine3.17 AS builder
WORKDIR /app
COPY . .

# Run stage
RUN go mod download
RUN go build -o main ./main.go

# Final stage
FROM alpine:3.17
COPY --from=builder /app/main .
COPY --from=builder /app/templates ./templates
# RUN chmod 755 main

# Expose port 8080 and start the application
EXPOSE 8080
CMD ["./main"]
