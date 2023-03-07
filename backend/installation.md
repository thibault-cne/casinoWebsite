# Backend setup

## Dependancies

You must have golang in version `1.20`.

## Environment

To run the app in dev mode you need to create a `development.env` file with the format of the `exemple.env` file.

- `MODE` : should be `dev` or `prod`.

## Start

You can use [task](https://taskfile.dev/installation/) to start the backend with :

```bash
cd casinoWebsite/backend
task run
```

Or manually with :

```bash
cd casinoWebsite/backend
docker-compose up -d
go run cmd/server/main.go
```
