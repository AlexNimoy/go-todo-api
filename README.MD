# todo api CA example

## .env

```
POSTGRES_USER=develop
POSTGRES_PASSWORD=develop
POSTGRES_PORT=5436
POSTGRES_HOST=localhost
SERVER_PORT=8085
```

## db migrate

install

`brew install golang-migrate`

`https://github.com/golang-migrate/migrate`

### usage

`migrate create -ext sql -dir ./schema -seq init`

## migrate up

`make migrate_up`

alias for

`migrate -path ./schema -database 'postgres://develop:develop@localhost:5436/develop?sslmode=disable' up`

## migrate down

`make migrate_down`

alias for

`migrate -path ./schema -database 'postgres://develop:develop@localhost:5436/develop?sslmode=disable' down`

## container psql console

`make db_cli`

alias for

`docker-compose exec postgres psql -U develop`

