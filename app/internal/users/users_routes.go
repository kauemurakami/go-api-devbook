package users

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetupUserRoutes(router *mux.Router) {
	router.HandleFunc("/users", CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users", GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{user}", GetUsersByNickOrName).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", GetUser).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", DeleteUser).Methods(http.MethodDelete)

}
