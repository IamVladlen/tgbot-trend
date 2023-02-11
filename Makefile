# Connect .env file
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

.PHONY:

db:
	docker compose up -d

run:
	go run cmd/main.go