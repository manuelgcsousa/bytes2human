BIN_NAME=b2h
BIN_DIR=bin
INSTALL_DIR=/usr/local/bin

build: $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BIN_NAME) cmd/main.go

$(BIN_DIR):
	mkdir -p $(BIN_DIR)

install: build
	sudo cp $(BIN_DIR)/$(BIN_NAME) $(INSTALL_DIR)

clean:
	go clean
	rm -rf $(BIN_DIR)
	sudo rm -f $(INSTALL_DIR)/$(BIN_NAME)
