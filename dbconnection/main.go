package main

import (
	"fmt"
	"golang/db"
	"golang/service"
	store "golang/userstore"
	"log"
	"net/http"
)

func main() {
	database, err := db.Connect()

	if err != nil {
		log.Println("database  failed ", err)
	}
	db.MigrateinPostgres(database)

	userStore := store.NewUserStore(database)
	userServices := service.NewUserServices(userStore)
	user, err := userServices.Register("rajan", "rajan@gmail.com", "hello@123")

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
