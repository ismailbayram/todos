# ToDo App
I have developed this application in order to improve myself and learn how to develop a web application in GoLang world.

# Running without docker
- Run the commands below after configuring the app.
- `go run cmd/migrate.go`
- `go run cmd/runserver.go`

# Running with docker
- `docker compose build todo-app`
- `docker compose up`
- `docker compose exec todos /app/migrate`

# Config
- The `example_config` file can be found in the `config` directory. You should copy of it with `config.yml` name in the same directory.

# API
- You can login with the request below;
```
curl --location --request POST 'http://localhost:8000/api/login/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "admin",
    "password": "123456"
}'
```
