package routes

import (
	"encoding/json"
	"golang/service"
	"golang/utils"
	"net/http"
)

type registerRequest struct {
	Name     string `json: "name"`
	Email    string `json: "email"`
	Password string `json: "password"`
}

func UserRouter(service *service.UserSerice) *http.ServeMux{

	userR := http.NewServeMux()

	userR.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		var req registerRequest

		err := json.NewDecoder(r.Body).Decode(&req)

		if err != nil {
			// w.WriteHeader(http.StatusBadRequest)
			// w.Header().Set("Content-Type" , "application/json")
			// w.Write([]byte(`this is the wrong json formate`))
			utils.Error(w, http.StatusBadRequest, err)
			return 
		}

		user, err := service.Register(req.Name, req.Email, req.Password)

		if err != nil {
			// w.WriteHeader(http.StatusBadRequest)
			// w.Header().Set("Content-Type" , "application/json")
			// json.NewEncoder(w).Encode(err)

			utils.Error(w, http.StatusBadRequest, err)
			return 
		}

		// w.WriteHeader(http.StatusCreated)
		// w.Header().Set("Content-Type", "application/json")
		// json.NewEncoder(w).Encode(user)
		utils.Response(w, http.StatusCreated, user)
	})

	return userR
}
