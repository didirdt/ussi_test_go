FROM golang:1.25.3

WORKDIR /app

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download


# Copy source code
COPY . .

# Build the application
RUN go build -o main .

# Expose port
EXPOSE 8080

# Run the application
CMD ["./main"]