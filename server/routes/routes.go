package routes

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/paulacgates/todo-app-golang/controllers"
)

func HandleRoutes() {

	r := mux.NewRouter()
	r.HandleFunc("/api/todos", controllers.ShowAllTodos).Methods("GET")
	r.HandleFunc("/api/todos", controllers.CreateTodo).Methods("POST")
	r.HandleFunc("/api/todos/{id}/done", controllers.UpdateDone).Methods("PATCH")
	c := handlers.AllowedOrigins([]string{"http://localhost:5173"})
	http.ListenAndServe(":8080", handlers.CORS(c)(r))

}
