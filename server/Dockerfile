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

# Install required Go dependencies
RUN go get github.com/gorilla/mux
RUN go get github.com/jinzhu/gorm
RUN go get github.com/jinzhu/gorm/dialects/sqlite

COPY . .

# Set PATH to include go bin
ENV PATH="/go/bin:${PATH}"
ENV GOCACHE=/app/tmp/.cache

RUN mkdir -p /app/tmp/.cache

CMD ["air"]
