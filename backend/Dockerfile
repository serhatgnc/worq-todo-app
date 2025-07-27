# Stage 1: Build Go app
FROM golang:1.24.5-alpine AS build

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/main.go

# Stage 2: Runtime
FROM alpine:latest

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy binary from build stage
COPY --from=build /app/main .

# Expose port
EXPOSE 8080

# Set environment variables
ENV MONGO_URI=mongodb://mongodb:27017
ENV PORT=8080

# Run the application
CMD ["./main"] 