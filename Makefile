
help:
	@echo "run - Run application"

run:
	@go run ./cmd/todo/main.go

build:
	@echo "Creating builds/build-$$(date +'%H:%M')"
	@go build cmd/crud/main.go
	@mkdir -p "builds${FILENAME}"
	@mv main "builds/build-$$(date +'%H:%M')"
	@echo "Created builds/build-$$(date +'%H:%M')"


build-and-run:
	@go build
	@./crud

vendor:
	go mod vendor
