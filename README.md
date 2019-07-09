# mysql-healthcheck
Rich healthcheck for MySQL.

## run with local mysql docker container
```
# run mysql container
docker run -it -d -e MYSQL_USER=user -e MYSQL_PASSWORD=password -e MYSQL_DATABASE=test -e MYSQL_ROOT_PASSWORD=passord -p 3306:3306 --name mysql mysql

# run main.go
go run main.go -user user -password password -database test -query "show databases;" -host 127.0.0.1

# healthcheck
curl http://localhost:8080
OK
```
