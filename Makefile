.PHONY: test run image container up down prune

test:
	@go test ./...

run:
	@go run cmd/main.go

image:
	@docker build -t vgmrs/goexpert-weather-api:latest -f ./Dockerfile .

container: image
	@docker run -p 8000:8000 vgmrs/goexpert-weather-api:latest

up:
	@docker compose up

down:
	@docker compose down

prune:
	@docker system prune -a
