package main

import (
  "context"
	"fmt"
	"io"
	"log"
  "os"
	"net/http"

  "github.com/jackc/pgx/v4"
  "github.com/jackc/pgtype"
)

const (
  host     = "localhost"
  port     = 5432
  user     = "docker"
  password = "docker"
  dbname   = "docker"
)
func exitHandler(res http.ResponseWriter, req *http.Request) {
  os.Exit(0)
}

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
func newArticleHandler(res http.ResponseWriter, req *http.Request) {
}
func articleHandler(res http.ResponseWriter, req *http.Request) {
  urlExample := "postgres://docker:docker@localhost:5432/docker"
  conn, err := pgx.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

  var id int64
	var title string
  var date pgtype.Date
  var body string
	err = conn.QueryRow(context.Background(), "select id, title, date, body from public.article where id = $1", 42).Scan(&id, &title, &date, &body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(id, title, date, body)


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
	fmt.Fprintf(w, "Go web app powered by Nine Publishing 12345678")
}
func main() {
  http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/exit", exitHandler)
  http.HandleFunc("/hello", helloHandler)
  http.HandleFunc("/articles", newArticleHandler)
  http.HandleFunc("/articles/", articleHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		return
	}
}
