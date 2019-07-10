# mysql-healthcheck
Rich healthcheck for MySQL.

- You can execute a SQL by accessing this server.

- SQL is specified at server startup.


## run on local
Run server in local with MySQL container.

Usage
```sh
go run main.go --help
  -address string
    	The address to listen on (default "localhost:8080")
  -database string
    	The database to connect to
  -host string
    	The database host (default "localhost")
  -password string
    	The database password
  -port string
    	The database port (default "3306")
  -query string
    	The test query
  -user string
    	The database user name
```

Example
```sh
# run mysql container
docker run -it -d -e MYSQL_USER=user -e MYSQL_PASSWORD=password -e MYSQL_DATABASE=test -e MYSQL_ROOT_PASSWORD=passord -p 3306:3306 --name mysql mysql

# run server
go run main.go -user user -password password -database test -query "show databases;" -host localhost

# healthcheck
curl http://localhost:8080
```


## run on docker
Run server in docker with MySQL container.

Create `docker-compose.yml`. (see [example/docker-compose.yml](/example/docker-compose.yml))
```yaml
---
version: '3'
services:
  healthcheck:
    image: reireias/mysql-healthcheck
    ports:
      - 8080:8080
    environment:
      MYSQL_USER: user
      MYSQL_PASSWORD: password
      MYSQL_DATABASE: test
      MYSQL_QUERY: 'show databases;'
      MYSQL_HOST: mysql
      MYSQL_PORT: 3306
      ADDRESS: 0.0.0.0:8080
    depends_on:
      - mysql
  mysql:
    image: mysql
    environment:
      MYSQL_USER: user
      MYSQL_PASSWORD: password
      MYSQL_DATABASE: test
      MYSQL_ROOT_PASSWORD: password
```

Run
```sh
# create containers
docker-compose up -d

# healthcheck
curl http://localhost:8080
```
