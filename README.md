# MyBookListBackend

Live demo: https://api.mybooklist.ndtai.me

## For dev
Get Go-migrate (for database migrations management)
``` bash
mkdir -p ./bin/migrate
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz | tar -xvz -C ./bin/migrate
chmod u+x ./bin/migrate/migrate
```

## Missing / unusable yet

| Task                    | Priority | Level  |
| ----------------------- | -------- | ------ |
| Integrate with frontend | high     | medium |
| Notification mechanism  | medium   | hard   |
| Views                   | low      | hard   |
| Monitoring              | low      | hard   |

## Partial done
- API design/docs
  - Still need to update and improve
- Database
  - Improving
  - Missing views & foreign key constraint
- Core API implement
