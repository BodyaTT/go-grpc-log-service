.SILENT:
include .env

build:
	go build -o ./bin/app ./cmd

up:
	docker-compose up

down:
	docker-compose down

run: build
	./bin/app