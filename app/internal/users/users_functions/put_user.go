package users_functions

import "net/http"

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create user"))

}
