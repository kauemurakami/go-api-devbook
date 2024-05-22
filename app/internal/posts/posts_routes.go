package posts //use same package name inside your package
import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetupPostsRoutes(router *mux.Router) {
	router.HandleFunc("/users/posts", CreatePost).Methods(http.MethodPost)
	router.HandleFunc("/users/posts", GetPosts).Methods(http.MethodGet)
	router.HandleFunc("/users/posts/{id}", GetPost).Methods(http.MethodGet)
	router.HandleFunc("/users/posts/{id}", DeletePost).Methods(http.MethodDelete)
	router.HandleFunc("/users/posts", UpdatePost).Methods(http.MethodPut)
	router.HandleFunc("/users/posts/{id}/like", LikePost).Methods(http.MethodPut)
	router.HandleFunc("/users/posts/{id}/unlike", UnlikePost).Methods(http.MethodPut)
}
