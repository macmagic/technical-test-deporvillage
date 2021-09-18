.PHONY: help
help::
	@printf "\033[33m Usage:\033[39m\n"
	@printf "  make COMMAND : Launch command\n"
	@printf "\033[33m Setup:\033[39m\n"
	@printf "  make install\n"
	@printf "  make run\n"

.PHONY: install
install:
	@echo "Installing packages..."
	@go mod download
	@go get github.com/matryer/moq@latest
	@go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.40.1

.PHONY: run
run:
	@go run cmd/main.go