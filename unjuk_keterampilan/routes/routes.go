package routes

import (
    "github.com/gorilla/mux"
    "unjuk_keterampilan/controllers"
)

// SetUserRoutes mengatur semua route pengguna
func SetUserRoutes(r *mux.Router, uc *controllers.UserController) {
    r.HandleFunc("/users", uc.GetUsers).Methods("GET")
    r.HandleFunc("/users/{id}", uc.GetUserByID).Methods("GET")
    // Anda bisa menambahkan endpoint lain di sini
}
