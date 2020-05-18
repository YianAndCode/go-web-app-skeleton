BIN_DIR=./bin
PROJECT_NAME=my-web-app

all: build

build: app

app:
	go build -o $(BIN_DIR)/$(PROJECT_NAME) .