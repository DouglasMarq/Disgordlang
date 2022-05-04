install:
	go mod

run: docker/build
	docker run -d -t -i -e ENVIRONMENT="dev" \
	-e TOKEN=$$DISCORDBOTTOKEN \
	-e GUILDID=$$DISCORDGUILDID \
	nitra-bot

docker/build:
	docker build -t nitra-bot .

test/unit:
	go test ./...

run/dev:
	ENVIRONMENT="dev" TOKEN=$$DISCORDBOTTOKEN GUILDID=$$DISCORDGUILDID go run cmd/main.go
