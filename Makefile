.PHONY: all build test clean

BIN := renderer
ALL_SRC := $(wildcard **/*.go)
SRC := $(filter-out %_test.go, $(wildcard **/*.go))
ENTRY := main.go

all: build

$(BIN): $(SRC)
	go build -o $(BIN) $(ENTRY)

build: $(BIN)

test: $(ALL_SRC)
	go test -v ./...

clean:
	rm -f $(BIN)

run: build
	./$(BIN)

.PHONY: all build test clean run
