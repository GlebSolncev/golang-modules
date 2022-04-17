help:
	@echo "run - Run application"

run:
	@go run .

build:
	@go build


build-and-run:
	@go build
	@./crud
