BINARY_NAME=CICDEX02

all: test build

build:
	go build -o $(BINARY_NAME) -v

test:
	go test -v

clean:
	rm -f $(BINARY_NAME)

run:
	go build -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
