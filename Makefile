# Connect .env file
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

.PHONY:

run:
	go run cmd/bot/main.go

db:
	docker compose -f docker-testdbs.yml up -d

build:
	docker compose up -d