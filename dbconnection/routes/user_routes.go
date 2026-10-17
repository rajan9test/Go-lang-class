package routes

import (
	"golang/handler"
	"net/http"
)

func UserRouter() *http.ServeMux {

	userR := http.NewServeMux()

	userR.HandleFunc("/register", handler.Register)

	return userR
}
