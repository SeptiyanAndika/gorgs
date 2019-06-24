## Description
Simple project used to learn and try clean arch, api gateway, and docker. this project have 2 services using golang run on container, and used kong api gateway for exposed to the public

This project has some domain layer :
 * Models Layer  : contains entities, dto, like schema table and or schema response object
 * Repository / Adapter Layer  : contains function to connect databse or another thir party
 * Usecase Layer  : contains function logic, use case or user stories
 * Logging Layer  : function for logging usecase , like a paramater, time process and error in usecase
 * Delivery Layer : function for expose usecase or logic to outside can be via http, grpc, lambda and others.


## Technology used:
- Golang
- Echo Framework (Rest API)
- Kong (Api Gateway)
- PostgreSql
- Docker


## Quick Setup
install docker first after that run docker compose
```
 docker-compose up --build
```

## Member Services
services get list member from github organisation, after that get detail each user and sorted by followers

## Comment Services
services crud comment to organisation name, data save on db postgresql 

## __http-test__
list and test endpoint for member and comments
