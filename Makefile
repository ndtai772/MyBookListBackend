DB_URL=postgresql://dev:123@localhost:5432/my_book_list?sslmode=disable
REAL_DB_URL=postgresql://dev:123@api.mybooklist.ndtai.me:5432/my_book_list?sslmode=disable
MIGRATE=./bin/migrate/migrate

dbup:
	docker-compose up -d
	@sleep 3
	$(MIGRATE) -path db/migrations -database "$(DB_URL)" -verbose up

dbdown:
	docker-compose down -v

dbreload: dbdown dbup

build:
	GOOS=linux GOARCH=amd64 go build -o bin/server

run:
	./bin/server
