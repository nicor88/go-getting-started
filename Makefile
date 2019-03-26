APP_NAME = "my-app"

run:
	go run -race ./cmd/$(APP_NAME)


build: clean
	go build -o bin/$(APP_NAME) ./cmd/$(APP_NAME)

clean:
	rm -rf bin/$(APP_NAME)

execute: build
	./bin/$(APP_NAME)
