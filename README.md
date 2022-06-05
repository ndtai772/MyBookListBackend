# MyBookListBackend

Live demo: https://api.mybooklist.ndtai.me

## For dev
Get Go-migrate (for database migrations management)
``` bash
mkdir -p ./bin/migrate
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz | tar -xvz -C ./bin/migrate
chmod u+x ./bin/migrate/migrate
```