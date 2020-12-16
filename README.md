# go_crud

A simple CRUD exercise using Golang with a MySQL Docker container as DB


## Prerequisites
- Golang
- Docker
- Docker-Compose
## Getting started

- Spin up the MySQL container:
```bash
$ docker-compose up -d
```

- Check if the MySQL container is up and running:
```bash
$ docker ps | grep -i devbook_db
f9264aa85e0e        mysql:5.7           "docker-entrypoint.s…"   21 hours ago        Up 7 hours          0.0.0.0:3306->3306/tcp, 33060/tcp   devbook_db
```

- Spin up the CRUD server:
```bash
$ ./crud
Server listening on port 5000
```

- Create the Users table which the API will make use of:
```sql
CREATE TABLE users(
	id int auto_increment primary key,
	name varchar(50) not null,
	email varchar(50) not null
) Engine=InnoDB;
```

<br />

## Issuing API calls

One can make use of Postman or Insomnia to hit the API endpoints for testing purposes.

<br />

### Some examples:

<br />

#### For creating a user:
- HTTP Method: `POST`
- URL: `localhost:5000/users`
- JSON payload:
```json
{
	"name": "Petter Griffin",
	"email": "birdistheword@gmail.com"
}
```

<br />

#### For fetching all users:
- HTTP Method: `GET`
- URL: `localhost:5000/users`

<br />

#### For fetching an specific user of ID "1":
- HTTP Method: `GET`
- URL: `localhost:5000/users/1`

<br />

#### For updating a user:
- HTTP Method: `PUT`
- URL: `localhost:5000/users/1`
- JSON payload:
```json
{
	"name": "Glenn Quagmire",
	"email": "guiggity@gmail.com"
}
```

<br />

#### For deleting a user of ID "1":
- HTTP Method: `DELETE`
- URL: `localhost:5000/users/1`