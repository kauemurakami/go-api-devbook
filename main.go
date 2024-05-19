package main

import (
	app_config "api-social-media/app/core/config"
	"api-social-media/app/core/db"
	"api-social-media/app/core/middlewares"
	"api-social-media/app/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	app_config.SetupEnvironments()

	db.SetupDB()
	fmt.Printf("Run API :%s", app_config.API_port)
	router := routes.SetupAppRoutes()
	router.Use(middlewares.SetupHeadersMiddleware)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", app_config.API_port), router))

}
