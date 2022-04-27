DB_URL=postgresql://dev:123@localhost:5432/my_book_list?sslmode=disable
MIGRATE=./bin/migrate/migrate

up:
	docker-compose up

createdb:
	docker exec backend-db-1 createdb --username=dev --owner=dev my_book_list

migrateup:
	$(MIGRATE) -path db/migrations -database "$(DB_URL)" -verbose up

build:
	go build -o bin/server main.go

run:
	./bin/server