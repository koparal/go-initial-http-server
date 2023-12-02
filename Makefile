.PHONY: build run test clean

APP_NAME = server
DOCKER_IMAGE_NAME = server-image
DOCKER_CONTAINER_NAME = server-container

build:
	go build -o $(APP_NAME) cmd/server/main.go

run: build
	./$(APP_NAME)

test:
	go test ./...

clean:
	go clean
	rm -f $(APP_NAME)

docker-build:
	docker build -t $(DOCKER_IMAGE_NAME) .

docker-run:
	docker run --rm --name $(DOCKER_CONTAINER_NAME) -p 8080:8080 $(DOCKER_IMAGE_NAME)

docker-stop:
	docker stop $(DOCKER_CONTAINER_NAME)
