package routes

import (
	"api-social-media/app/internal/users"

	"github.com/gorilla/mux"
)

func SetupAppRoutes() *mux.Router {
	// users.SetupUserRoutes(router)
	router := mux.NewRouter()
	users.SetupUserRoutes(router)
	return router
}
