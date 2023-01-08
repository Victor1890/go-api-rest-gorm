package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Victor1890/go-api-rest/src/database"
	"github.com/Victor1890/go-api-rest/src/models"
	"github.com/gorilla/mux"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	database.DB.Find(&users)

	json.NewEncoder(w).Encode(&users)
}
func GetOneUserHandler(w http.ResponseWriter, r *http.Request) {

	var params = mux.Vars(r)
	var user models.User

	var id = params["ida"]

	database.DB.Find(&user, id)

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	json.NewEncoder(w).Encode(&user)

}
func PostUsersHandler(w http.ResponseWriter, r *http.Request) {

	var payload models.User
	json.NewDecoder(r.Body).Decode(&payload)

	createdUser := database.DB.Create(&payload)
	userError := createdUser.Error

	if userError != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(userError.Error()))
		return
	}

	w.WriteHeader(http.StatusAccepted)

	json.NewEncoder(w).Encode(&payload)
	return

}
func DeleteUsersHandler(w http.ResponseWriter, r *http.Request) {

	var params = mux.Vars(r)
	var user models.User

	var id = params["ida"]

	database.DB.Find(&user, id)

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	database.DB.Delete(&user)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(200)

}
