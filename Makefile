DB_URL=mysql://dev:123@tcp(127.0.0.1:3307)/MyBookList
MIGRATE=./bin/migrate/migrate

migrateup:
	$(MIGRATE) -path db/migrations -database "$(DB_URL)" -verbose up

migrateup1:
	$(MIGRATE) -path db/migrations -database "$(DB_URL)" -verbose up 1

migratedown:
	$(MIGRATE) -path db/migrations -database "$(DB_URL)" -verbose down

migratedown1:
	$(MIGRATE) -path db/migrations -database "$(DB_URL)" -verbose down 1

getgomigrate:
	mkdir -p ./bin/migrate
	curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz | tar -xvz -C ./bin/migrate
	chmod u+x ./bin/migrate/migrate