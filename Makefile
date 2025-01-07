export PATH := $(PWD)/bin:$(PATH)

run: run-server
build: build-server

run-server:
	cd hyperlist/server && \
	./bin/main

build-server:
	cd hyperlist/server && \
	templ generate && \
	npx tailwindcss -c ./tailwind.config.js -i ./main.css -o ./render/public/tailwind.css && \
	go mod tidy && \
	go mod download && \
	go build -o bin/main main.go