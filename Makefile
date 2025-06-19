.PHONY: swag fmt check-swag

# Check if swag is installed, if not, install it
check-swag:
	@command -v swag >/dev/null 2>&1 || { \
		echo "swag not found, installing..."; \
		go install github.com/swaggo/swag/cmd/swag@latest; \
	}

# Generate Swagger docs
swag: check-swag
	$(HOME)/go/bin/swag init --dir ./cmd,./internal --output ./docs

# Format Go files
fmt:
	go fmt ./...
