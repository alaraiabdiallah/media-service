# Media Service

Small service to store your file

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

you need these software installed on your machine
- golang version 1.13
- IDE or editor (Goland or VSCode with GO extention)
- Docker engine v19.03.8
- Docker compose 1.25.4

### Initial Set-up

- Copy or clone this repository to $GOPATH/git.gits.id/RnD/WEB/golang-echo-starter
- run `docker-compose up -d --build` to run the program locally using docker compose

### Build
- to build the program run `go build cmd/main/server.go` to generate file `main` in current location
- to run the program `./main`

## Running the tests

In this project we use Testify to create the test. 

- To test the code, create file *_test.go and create a function with name that begin with "Test", for example "TestGetOneItem"
- To run the test type `go test -v`

## Folder Structure
```
media-service/
┣ app/
┣ data_sources/
┣ helpers/
┃ ┣ env.go
┃ ┣ http_responses.go
┃ ┗ validator.go
┣ models/
┣ protocols/
┃ ┗ http/
┃   ┣ handlers/
┃   ┣ middlewares/
┃   ┣ routers/
┃   ┗ http.go
┣ tests/
┃ ┣ common_test.go
┃ ┣ sender_test.go
┃ ┗ transport_test.go
┣ .gitignore
┣ Dockerfile
┣ README.md
┣ docker-compose.yaml
┣ go.mod
┣ go.sum
┗ main.go
```

| File        | Description | 
| :---        |    :----   | 
| `app/` | Main applications for this project. This is where all the logic located    |
| `data_sources/`   | Contain function that connect to mongodb |
| `helpers/*` | Contains something that have a possibility to be use over and over. For example in here we create response.go that we call when we give response    |
| `models/*`    | contain files that describe mongo's schema |
| `protocols/*` | In here we describe application route, middleware and which function from app that we need to call |
| `tests/` | unit test is in here |

## How to use
- to check the application run using docker compose
- list of environment variable needed 
```
MONGO_HOST=
MONGO_PORT=
MONGO_USER=
MONGO_PASS=
MONGO_DBNAME=
HTTP_PROT_PORT=
API_KEY=
```

## Deployment

Add additional notes about how to deploy this on a live system

## Built With

* [Echo](https://echo.labstack.com/guide) - The web framework used
* [mongo](https://godoc.org/go.mongodb.org/mongo-driver/mongo) - Orm used
* [Testify] (https://godoc.org/github.com/stretchr/testify) - Testing library

## Contributing

## Versioning
* v1.0.1 = edit readme
* v1.0.0 = initial version

## Authors
* [Ala Rai](https://github.com/alaraiabdiallah/) - The Original Author

## Acknowledgments

* Hat tip to anyone whose code was used
* Inspiration
* etc
