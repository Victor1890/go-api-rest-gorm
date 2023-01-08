package routes

import (
	"github.com/Victor1890/go-api-rest/src/controllers"
	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) *mux.Router {

	r.HandleFunc("/users", controllers.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", controllers.GetOneUserHandler).Methods("GET")
	r.HandleFunc("/users", controllers.PostUsersHandler).Methods("POST")
	r.HandleFunc("/users/{id}", controllers.DeleteUsersHandler).Methods("DELETE")

	return r

}
