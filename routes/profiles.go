package routes

import (
	"wetalk/handlers"
	"wetalk/pkg/middleware"
	"wetalk/pkg/mysql"
	"wetalk/repositories"

	"github.com/gorilla/mux"
)

func ProfileRoutes(r *mux.Router) {
	ProfileRepository := repositories.RepositoryProfile(mysql.DB)
	h := handlers.HandlerProfile(ProfileRepository)

	r.HandleFunc("/profiles", middleware.Auth(h.FindProfiles)).Methods("GET")
	r.HandleFunc("/profile", middleware.Auth(h.FindProfile)).Methods("GET")
	r.HandleFunc("/profile/{id}", middleware.Auth(h.GetProfile)).Methods("GET")
	r.HandleFunc("/profile", middleware.Auth(middleware.UploadFile(h.CreateProfile))).Methods("POST")
	r.HandleFunc("/profile/{id}", middleware.Auth(middleware.UploadFile(h.UpdateProfile))).Methods("PATCH")
	r.HandleFunc("/profile/{id}", middleware.Auth(h.DeleteProfile)).Methods("DELETE")
}
