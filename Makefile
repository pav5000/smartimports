
BIN_DIR=$(CURDIR)/bin

build:
	go build -o ${BIN_DIR}/smartimports github.com/pav5000/smartimports/cmd/smartimports
