package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/paulacgates/todo-app-golang/database"
	"github.com/paulacgates/todo-app-golang/models"
)

func ShowAllTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var p []models.Todo
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var t models.Todo
	json.NewDecoder(r.Body).Decode(&t)
	database.DB.Create(&t)
	json.NewEncoder(w).Encode(t)
}

func UpdateDone(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var t models.Todo
	database.DB.First(&t, id)
	t.Done = true
	database.DB.Save(&t)
	json.NewEncoder(w).Encode(t)
}
