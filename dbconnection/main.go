package main

import (
	"encoding/json"
	"fmt"
	"golang/db"
	"golang/service"
	store "golang/userstore"
	"log"
	"net/http"
)

type registerRequest struct {
	Name     string `json: "name"`
	Email    string `json: "email"`
	Password string `json: "password"`
}

func main() {
	database, err := db.Connect()

	if err != nil {
		log.Println("database  failed ", err)
	}
	db.MigrateinPostgres(database)

	userStore := store.NewUserStore(database)
	userServices := service.NewUserServices(userStore)

	router := http.NewServeMux()

	router.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {

		var req registerRequest

		err := json.NewDecoder(r.Body).Decode(&req)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type" , "application/json")
			w.Write([]byte(
				(`hello the json formate are not sutable `)
			))
		}

	})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)

	if err != nil {
		fmt.Println(err)
	}
	server := &http.Server{
		Addr: ":8080",
	}

	log.Println("server is running on the port of :8080")

	err = server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
