build:
	@go build -o bin/toronto-wheels ./cmd/toronto-wheels

run: build
	@./bin/toronto-wheels

test:
	@go test -v -race ./...

docker-up:
	@docker compose up -d --build

clean:
	rm -rf ./bin
