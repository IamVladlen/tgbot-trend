# Connect .env file
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

.PHONY:

run:
	go run cmd/bot/main.go

build:
	docker compose up -d