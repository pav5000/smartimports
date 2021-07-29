
BIN_DIR=$(CURDIR)/bin
PACKAGE=github.com/pav5000/smartimports/cmd/smartimports

build:
	go build -o ${BIN_DIR}/smartimports ${PACKAGE}

run:
	go run ${PACKAGE}

test:
	go test ./...
