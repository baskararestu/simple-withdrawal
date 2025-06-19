# Build stage
FROM golang:1.24.4 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/main.go

# Final image
FROM alpine:latest  
WORKDIR /root/
COPY --from=builder /app/app .
COPY ./docs ./docs

EXPOSE 6005

CMD ["./app"]
