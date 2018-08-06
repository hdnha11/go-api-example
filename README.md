# Go Rest API example

## Development

### Using Docker

#### Get Docker

Get [Docker Community Edition](https://www.docker.com/community-edition) for your platform *(Linux, Mac, Windows)*.

#### Spin up the API server and Database

```
$ docker-compose up -d
```

Run this command to verify:

```
$ curl -i localhost:8080/status
```

It should output something like this:

```
HTTP/1.1 200 OK
Date: Wed, 13 Jun 2018 05:45:36 GMT
ontent-Length: 0
```

#### Stop the server

```
$ docker-compose down
```

#### Rebuild image

```
$ docker-compose up -d --build
```

#### Access to running containers

```
$ docker-compose exec api sh
$ docker-compose exec mongo-db mongo -u root -p password --authenticationDatabase "admin"
```

Or

```
$ docker container exec -it go-api-example_api_1 sh
$ docker container exec -it go-api-example_mongo-db_1 mongo -u root -p password --authenticationDatabase "admin"
```

#### View logs

```
$ docker-compose logs api
$ docker-compose logs mongo-db
```

Or

```
$ docker container logs go-api-example_api_1
$ docker container logs go-api-example_mongo-db_1
```

### Using your own local environment

#### Install Go

[Download and Install Go](https://golang.org/doc/install)

#### Setup Go Workspace

Create your Go [workspace](https://golang.org/doc/code.html#Workspaces)

#### Install MongoDB

Using Docker

```
$ docker container run --name mongodb -d \
  -e MONGO_INITDB_ROOT_USERNAME=root \
  -e MONGO_INITDB_ROOT_PASSWORD=password \
  -p 27017:27017 -v mongo-data:/data/db \
  mongo:latest
```

Or [install](https://docs.mongodb.com/manual/installation/) on your local machine

#### Set environment variables

Clone `.env.sample` and named it `.env`

#### Build and Run

```
$ go run main.go
```

Or

```
$ go build
$ ./go-api-example
```

Or

```
$ go install
$ go-api-example
```

## API

### Projects

| Resource                      | Method | Description                       |
|:------------------------------|:-------|:----------------------------------|
| `/api/v1/projects`            | GET    | Get all projects                  |
| `/api/v1/projects/:id`        | GET    | Get project by id                 |
| `/api/v1/projects?name=<name>`| GET    | Find project by name              |
| `/api/v1/projects`            | POST   | Create project                    |
| `/api/v1/projects/:id`        | PUT    | Update project                    |
| `/api/v1/projects/:id`        | DELETE | Delete project                    |

## Troubleshooting

### Remove MongoDB data volume

List current volumes:

```
$ docker volume ls
```

Remove data volume:

```
$ docker volume rm go-api-example_mongo-data
```

Or

```
$ docker-compose down --volumes
```
