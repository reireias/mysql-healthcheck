package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var (
	user     = flag.String("user", "", "The database user name")
	password = flag.String("password", "", "The database password")
	host     = flag.String("host", "localhost", "The database host")
	port     = flag.String("port", "3306", "The database port")
	database = flag.String("database", "", "The database to connect to")
	query    = flag.String("query", "", "The test query")
	address  = flag.String("address", "localhost:8080", "The address to listen on")
)

func main() {
	flag.Parse()
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", *user, *password, *host, *port, *database))
	if err != nil {
		fmt.Printf("Error opening database: %v\n", err)
	}

	// simple query exec handler
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		_, err := db.Exec(*query)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			res.Write([]byte(err.Error()))
			return
		}
		res.WriteHeader(http.StatusOK)
		res.Write([]byte("OK"))
		return
	})

	http.ListenAndServe(*address, nil)
}
