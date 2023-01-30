lint:
	golangci-lint run -c ./golangci.yml ./...

.PHONY: docs
docs:
	@echo "Generating docs..."
	swag fmt
	swag init -g purebackend.go

run:
	go run purebackend.go

build:
	go build -o purebackend.exe purebackend.go