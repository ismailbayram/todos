FROM golang:1.19-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./src ./src
COPY ./cmd ./cmd
COPY ./config ./config

RUN go build -o ./todos-app ./cmd/runserver.go
RUN go build -o ./migrate ./cmd/migrate.go

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/config /app/config
COPY --from=build /app/todos-app /app/todos-app
COPY --from=build /app/migrate /app/migrate

ENTRYPOINT ["./todos-app"]