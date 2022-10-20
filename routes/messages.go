package routes

import (
	"wetalk/handlers"
	"wetalk/pkg/mysql"
	"wetalk/repositories"

	"github.com/gorilla/mux"
)

func MessageRoutes(r *mux.Router) {
	MessageRepository := repositories.RepositoryMessage(mysql.DB)
	h := handlers.HandlerMessage(MessageRepository)

	r.HandleFunc("/messages", h.FindMessages).Methods("GET")
	r.HandleFunc("/message/{id}", h.GetMessage).Methods("GET")
	r.HandleFunc("/message", h.CreateMessage).Methods("POST")
	r.HandleFunc("/message/{id}", h.DeleteMessage).Methods("DELETE")
}
