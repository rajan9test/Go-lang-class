package main

import (
	"database/sql"
	"golang/routes"
	"log"
	"net/http"
)

var Database *sql.DB

func main() {

	router := routes.UserRouter()
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("server is running on the port of :8080")

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
