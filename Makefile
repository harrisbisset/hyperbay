run: run-server
build: build-server

run-server:
	cd cmd && \
	./bin/main

build-server:
	cd cmd && \
	go mod tidy && \
	go mod download && \
	go build -o bin/main main.go