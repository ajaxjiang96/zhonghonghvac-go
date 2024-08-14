cli:
	go build -o bin/cli cmd/*.go

clean:
	rm -rf bin

all: cli