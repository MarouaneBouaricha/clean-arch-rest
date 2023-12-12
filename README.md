# Clean Architecture with Go
Simple REST API that follows the Clean Architecture Principles
## Principles
- [x] Independent of Frameworks
- [x] Independent of Databases
- [x] Independent of UI
- [x] Independent of any external agency
- [x] Testable

## Running Postgres Instance
```
docker run --name postgres-database -p 5432:5432 -e POSTGRES_USER=tux -e POSTGRES_PASSWORD=mysecretpassword -e POSTGRES_DB=blog -d postgres
```

## Running the api
```
go run .
```

## Running tests for post-service
```
go test service/post-service.go service/post-service_test.go  -v
```