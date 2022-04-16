DB_URL=postgresql://dev:123@localhost:5432/my_book_list?sslmode=disable
MIGRATE=./bin/migrate/migrate

up:
	docker-compose up
	docker exec backend-db-1 createdb --username=dev --owner=dev my_book_list

migrateup:
	$(MIGRATE) -path db/migrations -database "$(DB_URL)" -verbose up

migrateup1:
	$(MIGRATE) -path db/migrations -database "$(DB_URL)" -verbose up 1

migratedown:
	$(MIGRATE) -path db/migrations -database "$(DB_URL)" -verbose down

migratedown1:
	$(MIGRATE) -path db/migrations -database "$(DB_URL)" -verbose down 1
