package users

import (
	functions "api-social-media/app/internal/users/users_functions"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	functions.CreateUser(w, r)
}
func GetUsers(w http.ResponseWriter, r *http.Request) {
	functions.GetUsers(w, r)
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	functions.GetUser(w, r)
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	functions.UpdateUser(w, r)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	functions.DeleteUser(w, r)
}
