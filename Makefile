run: run-server
build: build-server

run-server:
	cd hyperlist/server && \
	./bin/main

build-server:
	cd hyperlist/server && \
	go mod tidy && \
	go mod download && \
	go build -o bin/main main.go