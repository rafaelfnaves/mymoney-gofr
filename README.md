# MyMoney

MyMoney is a personal finance management application built with Go, designed to help users track daily expenses.

## Description

This project leverages the [gofr.dev](https://gofr.dev) opinionated framework For accelerated microservice development. MyMoney aims to simplify personal finance tracking expenses.

## Technologies Used

- **Go**: Main programming language for backend development.
- **gofr.dev**: Framework for building scalable Go microservices.
- **PostgreSQL**: (Optional) For persistent data storage.
- **Docker**: (Optional) For containerized deployment.

## Getting Started

1. Clone the repository.
2. Install Go and dependencies of the project `go mod tidy`.
3. Execute `docker-compose up -d` to up Postgres database.
4. Run application with `go run cmd/main.go`.
5. Use extension [Rest Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) in VS Code to make requests and test API at file `api.http`. (Or other alternatives)

## License

MIT License