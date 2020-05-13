BIN=./bin/fvv

all: update build install

update:
	go get -u ./...
	go mod tidy
	go mod vendor

build:
	go build -v -o $(BIN) .

run: build
	$(BIN)

install:
	go install -v .