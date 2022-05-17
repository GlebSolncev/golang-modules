MAIN_PATH = "./cmd/todo"
BUILD_FILENAME=$$(date +'%H:%M:%S')
BUILD_PATH="builds/build-"$$(date +'%H:%M')

TEST=ON
ifeq ($(package),)
    STATUS='echo "Sorry but i have not the package. Add option package, ex: make setup package=[http]"'
else
    STATUS="go get $(package)"
endif

parser-run, pr:
	@go run cmd/parserURL/main.go

help:
	@echo "-----------------------------golang-modules-LIST-TODO-----------------------------"
	@echo "\t up\t\t\t\t- Build and run application"
	@echo "\t run\t\t\t\t- Run application"
	@echo "\t generate(gen)\t\t\t- Generate list packages"
	@echo "\t build\t\t\t\t- Build app. save in folder: 'builds'"
	@echo "\t build-and-run(bar)\t\t- Build app in 'builds' and run"
	@echo "\t get [package=[github....]]\t- install external package"
	@echo "\t test\t\t\t\t- Testing application"
	@echo "\t run\t\t\t\t- Run application"
	@echo "\t swagger-setup\t\t\t- Init swagger"
	@echo "------------------------------------------------------------------------"

godoc:
	godoc --http :6060

# (dev command) to create model
create-model:
	ent init  --target pkg/ent/schema $(model)


swagger-setup, si:
	@swag init -g ./cmd/todo/main.go --output docs/
	@echo "OK"

up: generate, swagger-setup, enumer-setup, run

enumer-setup, es:
	enumer  --type=TodoStatus -json internal/app/models/todo.go

generate, gen:
	go generate ./pkg/ent/generate.go
	go generate ./internal/models/todo/todo.go

run:
	@go run $(MAIN_PATH)

build:
	@go build -o $(BUILD_PATH) $(MAIN_PATH)
	@echo "OK. DONE! -> $(BUILD_PATH)"

build-and-run, bar: build
	@./$(BUILD_PATH)

get: setup

test:
	@go test $(MAIN_PATH)

setup:
	@eval $(STATUS)