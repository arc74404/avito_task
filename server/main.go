package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// "fmt"
// "net/http"
// "github.com/gorilla/mux"

func main() {
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=avito_db sslmode=disable")

	defer db.Close()

	if err != nil {
		log.Print("failed open db")
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	if err = db.Ping(); err != nil {
		log.Fatal("Bd does not answer:", err)
	}
}
