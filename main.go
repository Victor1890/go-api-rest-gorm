package main

import (
	"fmt"
	"net/http"

	"github.com/Victor1890/go-api-rest/src/database"
	"github.com/Victor1890/go-api-rest/src/models"
	"github.com/Victor1890/go-api-rest/src/routes"
	"github.com/gorilla/mux"
)

func main() {

	database.DbConnection()

	database.DB.AutoMigrate(models.User{})
	database.DB.AutoMigrate(models.Task{})

	var router = mux.NewRouter()

	router.HandleFunc("/", routes.HomeHandler)

	router = routes.UserRoutes(router)

	fmt.Println("http://localhost:3000")
	http.ListenAndServe(":3000", router)
}
