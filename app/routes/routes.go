package routes

import (
	"api-social-media/app/internal/auth"
	"api-social-media/app/internal/followers"
	"api-social-media/app/internal/posts"
	"api-social-media/app/internal/users"

	"github.com/gorilla/mux"
)

func SetupAppRoutes() *mux.Router {
	router := mux.NewRouter()
	auth.SetupAuthRoutes(router)
	users.SetupUserRoutes(router)
	followers.SetupFollowersRoutes(router)
	posts.SetupPostsRoutes(router)
	return router
}
