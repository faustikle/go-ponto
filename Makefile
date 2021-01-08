CMD_PKG=./cmd/ponto
BINARY_NAME=ponto

build: clean build-unix build-windows
.PHONY: build

build-unix: clean
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -v -o $(BINARY_NAME) $(CMD_PKG)
.PHONY: build-unix

build-windows: clean
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -v -o $(BINARY_NAME).exe $(CMD_PKG)
.PHONY: build-windows

clean:
	rm -f ./$(BINARY_NAME)
	rm -f ./$(BINARY_NAME).exe
.PHONY: clean
