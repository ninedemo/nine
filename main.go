package main

import (
	"fmt"
//  "database/sql"
	"io"
	"log"
	"net/http"

//  _ "github.com/lib/pq"
)

const (
  host     = "localhost"
  port     = 5432
  user     = "docker"
  password = "docker"
  dbname   = "postgres"
)

func helloHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<doctype html>
<html>
	<head>
		<title>Hello Gopher</title>
	</head>
	<body>
		Hello Gopher </br>
	</body>
</html>`,
	)
}
func dataHandler(res http.ResponseWriter, req *http.Request) {
  /*
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
      "password=%s dbname=%s sslmode=disable",
      host, port, user, password, dbname)
    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
      panic(err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
      panic(err)
    }

    fmt.Println("Successfully connected!")
*/
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<doctype html>
<html>
	<head>
		<title>Hello Data</title>
	</head>
	<body>
		Hello Data </br>
	</body>
</html>`,
	)
}
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go web app powered by Nine Publishing 11")
}
func main() {

	http.HandleFunc("/", defaultHandler)
  http.HandleFunc("/hello", helloHandler)
  http.HandleFunc("/data", dataHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		return
	}
}
