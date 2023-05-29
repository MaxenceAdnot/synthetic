BINARY_NAME=synthetic

.PHONY: build
build:
	mkdir -p ./bin
	go build -o ./bin/${BINARY_NAME} ./cmd/api/main.go

.PHONY: build-dist
build-dist:
	mkdir -p ./bin
	GOARCH=arm64 GOOS=darwin go build -o ./bin/${BINARY_NAME}-arm64-darwin ./cmd/api/main.go
	GOARCH=arm64 GOOS=linux go build -o ./bin/${BINARY_NAME}-arm64-linux ./cmd/api/main.go
	GOARCH=arm64 GOOS=windows go build -o ./bin/${BINARY_NAME}-arm64-windows ./cmd/api/main.go
	GOARCH=amd64 GOOS=darwin go build -o ./bin/${BINARY_NAME}-amd64-darwin ./cmd/api/main.go
	GOARCH=amd64 GOOS=linux go build -o ./bin/${BINARY_NAME}-amd64-linux ./cmd/api/main.go
	GOARCH=amd64 GOOS=windows go build -o ./bin/${BINARY_NAME}-amd64-windows ./cmd/api/main.go

.PHONY: run
run: build
	./bin/${BINARY_NAME}

.PHONY: clean
clean:
	go clean
	rm -rf ./bin/

.PHONY: test
test:
	go test -race ./...

.PHONY: test_coverage
test_coverage:
	go test ./... -coverprofile=coverage.out

.PHONY: dep
dep:
	go mod download

.PHONY: vet
vet:
	go vet ./...

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: container-build
container-build:
	podman build -t ${BINARY_NAME} .

.PHONY: container-run
container-run: container-build
	podman run -p 8080:8080 ${BINARY_NAME}