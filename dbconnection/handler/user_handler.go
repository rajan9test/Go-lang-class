package handler

import (
	"encoding/json"
	"fmt"
	"golang/db"
	"golang/service"
	store "golang/userstore"
	"golang/utils"
	"log"
	"net/http"
)

type registerRequest struct {
	Name     string `json: "name"`
	Email    string `json: "email"`
	Password string `json: "password"`
}

func Register(w http.ResponseWriter, r *http.Request) {

	var req registerRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		// w.WriteHeader(http.StatusBadRequest)
		// w.Header().Set("Content-Type" , "application/json")
		// w.Write([]byte(`this is the wrong json formate`))
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	Database, err := db.Connect()

	if err != nil {
		log.Println("database  failed ", err)
	}

	userStore := store.NewUserStore(Database)
	userServices := service.NewUserServices(userStore)

	user, err := userServices.Register(req.Name, req.Email, req.Password)

	if err != nil {
		// w.WriteHeader(http.StatusBadRequest)
		// w.Header().Set("Content-Type" , "application/json")
		// json.NewEncoder(w).Encode(err)

		fmt.Println(err)
		utils.Error(w, http.StatusForbidden, err.Error())
		return
	}

	// w.WriteHeader(http.StatusCreated)
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(user)
	utils.Response(w, http.StatusCreated, user)
}
