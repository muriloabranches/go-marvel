APP_NAME=go-marvel

default: test run
run: build execute clean

build:
	go build -o bin/${APP_NAME} ./cmd

execute:
	./bin/${APP_NAME}

test: 
	go test -v --cover ./...

clean:
	rm -rf bin