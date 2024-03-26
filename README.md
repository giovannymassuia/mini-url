# Mini URL in Go

## Description

This is a simple URL shortener written in Go. It uses an in-memory map to store the URLs and their shortened versions.
The shortened version is generated using just the id of the URL in the map.

## How to run

Using Makefile:

API:

- `make run-api`

Docs:

- `make run-docs`

or

Go API:

- `cd app`
- `go run main.go`

Docusaurus:

- `cd docs`
- `npm run start`

## Swagger docs

Swaager UI:

- `http://localhost:8080/swagger/index.html`

Swagger JSON:

- `http://localhost:8080/swagger/doc.json`

Regenerate swagger docs:

- `swag init`

After generating the docs, copy the `app/docs/swagger.json` file to the `docs` (docusaurus) folder.

## swagger setup

- https://github.com/swaggo/swag

- https://github.com/swaggo/http-swagger

- `go install github.com/swaggo/swag/cmd/swag@latest`

- `swag init` in the root of the project where the `main.go` file is located
