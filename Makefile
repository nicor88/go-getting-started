APP_NAME = "my-app"

run:
	go run -race ./cmd/$(APP_NAME)


build: clean
	go build ./cmd/$(APP_NAME)

clean:
	rm -rf $(APP_NAME)

execute: build
	./$(APP_NAME)
