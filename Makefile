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

make-migrate:
	migrate create -ext sql -dir ./scheduler-service/migrations -seq init

migrate:
	migrate -path ./scheduler-service/migrations/ -database 'postgres://testUser:testPass@localhost:5436/testDB?sslmode=disable' up