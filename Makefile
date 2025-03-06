BINARY_NAME=pack-distributor
CMD_PATH=./cmd/api
DOCKER_IMAGE=cvele/re-partners
CONTAINER_NAME=pack-distributor-test

GO_FLAGS=-v
GO_FILES=$(shell find . -name '*.go' -not -path "./vendor/*")

.PHONY: all
all: fmt vet build

.PHONY: build
build:
	@echo "Building the binary..."
	mkdir -p bin
	@if [ "$(shell uname)" = "Darwin" ]; then \
		GOOS=linux CGO_ENABLED=0 go build -o bin/$(BINARY_NAME) $(CMD_PATH)/main.go; \
	else \
		GOOS=linux CGO_ENABLED=1 go build -o bin/$(BINARY_NAME) $(CMD_PATH)/main.go; \
	fi

.PHONY: run
run:
	@echo "Running the application..."
	go run $(CMD_PATH)/main.go

.PHONY: test
test:
	@echo "Running unit tests..."
	go test ./... -cover -v

.PHONY: fmt
fmt:
	@echo "Formatting code..."
	go fmt ./...

.PHONY: vet
vet:
	@echo "Linting code..."
	go vet ./...

.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm -rf bin

.PHONY: docker-build
docker-build:
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE) .

.PHONY: docker-run
docker-run:
	@echo "Running Docker container..."
	docker run --rm -p 8080:8080 $(DOCKER_IMAGE)

.PHONY: docker-push
docker-push:
	@echo "Pushing Docker image to registry..."
	docker push $(DOCKER_IMAGE)

.PHONY: docker-stop
docker-stop:
	@echo "Stopping Docker container..."
	-docker stop $(CONTAINER_NAME)

.PHONY: prod-test
prod-test: 
	@echo "Running curl tests..."

	@echo "GET /api/v1/packs"
	curl -i https://repartnerstest.fly.dev/api/v1/packs
	@echo "\n"

	@echo "POST /api/v1/packs"
	curl -i -X POST https://repartnerstest.fly.dev/api/v1/packs \
		-H "Content-Type: application/json" \
		-d '{"size":750}'
	@echo "\n"

	@echo "GET /api/v1/calculate?order=12001"
	curl -i "https://repartnerstest.fly.dev/api/v1/calculate?order=12001"
	@echo "\n"

	@echo "DELETE /api/v1/packs/1"
	curl -i -X DELETE "https://repartnerstest.fly.dev/api/v1/packs/1"
	@echo "\n"

.PHONY: docker-test
docker-test: docker-build
	@echo "Running Docker container in detached mode..."
	docker run -d --rm --name $(CONTAINER_NAME) -p 8080:8080 $(DOCKER_IMAGE)
	@echo "Waiting for the container to initialize..."
	sleep 5
	@echo "Running curl tests..."

	@echo "GET /api/v1/packs"
	curl -i http://localhost:8080/api/v1/packs
	@echo "\n"

	@echo "POST /api/v1/packs"
	curl -i -X POST http://localhost:8080/api/v1/packs \
		-H "Content-Type: application/json" \
		-d '{"size":750}'
	@echo "\n"

	@echo "GET /api/v1/calculate?order=12001"
	curl -i "http://localhost:8080/api/v1/calculate?order=12001"
	@echo "\n"

	@echo "DELETE /api/v1/packs/1"
	curl -i -X DELETE "http://localhost:8080/api/v1/packs/1"
	@echo "\n"

	@echo "Stopping Docker container..."
	docker stop $(CONTAINER_NAME)

.PHONY: tidy
tidy:
	@echo "Tidying Go modules..."
	go mod tidy

.PHONY: deps
deps:
	@echo "Installing dependencies..."
	go mod download

.PHONY: rebuild
rebuild: clean tidy build

.PHONY: help
help:
	@echo "Available commands:"
	@echo "  make build         - Build the Go application"
	@echo "  make run           - Run the application"
	@echo "  make test          - Run unit tests"
	@echo "  make fmt           - Format code"
	@echo "  make vet           - Lint code"
	@echo "  make clean         - Clean build artifacts"
	@echo "  make docker-build  - Build Docker image"
	@echo "  make docker-run    - Run Docker container"
	@echo "  make docker-stop   - Stop Docker container"
	@echo "  make docker-test   - Build, run container, execute API tests, stop container"
	@echo "  make docker-push   - Push Docker image to registry"
	@echo "  make tidy          - Tidy up Go modules"
	@echo "  make deps          - Install dependencies"
	@echo "  make rebuild       - Clean, tidy, and build"
	@echo "  make prod-test     - Run curl tests against the production server"
	@echo "  make help          - Show this help message"
