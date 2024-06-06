.PHONY: build clean

default: build

build:
	go build -o go-flingo cmd/main.go

clean:
	rm -f reading-app
