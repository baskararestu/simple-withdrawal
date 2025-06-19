.PHONY: swag fmt

# Generate Swagger docs
swag:
	swag init --dir ./cmd,./internal --output ./docs

# Format Go files
fmt:
	go fmt ./...
