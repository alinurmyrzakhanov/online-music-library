.PHONY: build run migrate-up migrate-down clean

build:
	docker-compose build

run:
	docker-compose up

migrate-up:
	docker-compose run --rm migrate

migrate-down:
	docker-compose run --rm migrate -path=/migrations -database postgres://user:password@db:5432/music_library?sslmode=disable down

clean:
	docker-compose down --volumes --remove-orphans
