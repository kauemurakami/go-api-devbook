package users_functions

import "net/http"

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Recovery user"))

}
