APP_NAME = qrgenerator
BIN_DIR = bin
SRC = main.go

build:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(APP_NAME) $(SRC)

run: build
	./$(BIN_DIR)/$(APP_NAME)

clean:
	rm -rf $(BIN_DIR)
