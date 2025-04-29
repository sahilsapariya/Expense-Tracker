# Use the official Golang image as the builder
FROM golang:1.21-alpine

# Install dependencies for air
RUN apk add --no-cache git curl

# Install air binary
RUN go install github.com/cosmtrek/air@v1.40.4

# Set working directory inside container
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# Set PATH to include go bin
ENV PATH="/go/bin:${PATH}"
ENV GOCACHE=/app/tmp/.cache

RUN mkdir -p /app/tmp/.cache

CMD ["air"]
