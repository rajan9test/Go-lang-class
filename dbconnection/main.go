package main

import (
	"fmt"
	"golang/db"
	"golang/model"
	"golang/userstore"
	"log"
	"net/http"
)

func main() {
	database , err := db.Connect()

	if err != nil {
		log.Println("database  failed ", err)
	}
	db.MigrateinPostgres(database)

	userStore := userstore.NewUserStore(database);

	atul := model.User{
		Name: "Atul",
		Email: "atul@gmail.com",
		Password_hash: "Hello@123",
	}
	fmt.Println(atul)
     
	atul, err = userStore.Create(atul)

	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(atul)
	server := &http.Server{
		Addr: ":8080",
	}

	log.Println("server is running on the port of :8080")

	err = server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
