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
	GOOS=linux GOARCH=amd64 go build -o bin/server

run:
	./bin/server

deploy:
	cd ansible && ansible-playbook deploy.yaml -vv

svup:
	cd ansible && ansible-playbook base/up.yaml -vv
	cd ansible && ansible-playbook playbook.yaml -vv

svdown:
	cd ansible && ansible-playbook base/down.yaml -vv
