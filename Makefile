DB_URL=postgresql://dev:123@localhost:5432/my_book_list?sslmode=disable
MIGRATE=./bin/migrate/migrate

dbup:
	docker-compose up -d
	@sleep 3
	$(MIGRATE) -path db/migrations -database "$(DB_URL)" -verbose up

dbdown:
	docker-compose down -v

dbreload: dbdown dbup

build:
	go build -o bin/server main.go

run:
	./bin/server
