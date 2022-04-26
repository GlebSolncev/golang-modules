MAIN_PATH = "./cmd/todo"
BUILD_FILENAME=$$(date +'%H:%M:%S')
BUILD_PATH="builds/build-"$$(date +'%H:%M')

TEST=ON
ifeq ($(package),)
    STATUS='echo "Sorry but i have not the package. Add option package, ex: make setup package=[http]"'
else
    STATUS="go get $(package)"
endif

help:
	@echo "------------------------------------------------------------------------"
	@echo "\t run - Run application"
	@echo "\t build - Build app. save in folder: 'builds'"
	@echo "\t build-and-run - Build app in 'builds' and run"
	@echo "\t get [package=[github....]] - install external package"
	@echo "\t test - Testing application"
	@echo "\t run - Run application"
	@echo "------------------------------------------------------------------------"

generate:
	go generate ./internal/models/todo/todo.go
	go generate ./env/generate.go

run:
	@go run $(MAIN_PATH)

build:
	@go build -o $(BUILD_PATH) $(MAIN_PATH)
	@echo "OK. DONE! -> $(BUILD_PATH)"

build-and-run: build
	@./$(BUILD_PATH)

get: setup vend

test:
	@go test $(MAIN_PATH)

setup:
	@eval $(STATUS)

vend:
	go mod vendor
	@echo "OK"
