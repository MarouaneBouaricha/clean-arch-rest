# Clean Architecture with Go
Simple REST API that follows the Clean Architecture Principles
## Principles
- [x] Independent of Frameworks
- [x] Independent of Databases
- [x] Independent of UI
- [ ] Independent of any external agency
- [ ] Testable

## Running Postgres Instance
```
docker run --name postgres-database -p 5432:5432 -e POSTGRES_USER=tux -e POSTGRES_PASSWORD=mysecretpassword -e POSTGRES_DB=blog -d postgres
```

## Running the api
```
go run .
```