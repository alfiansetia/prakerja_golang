package main

import (
    "log"
    "net/http"
    "unjuk_keterampilan/controllers"
    "unjuk_keterampilan/config" // Import package config
    "unjuk_keterampilan/routes"
    "github.com/gorilla/mux"
)

func main() {
    // Dapatkan koneksi database dari package config
    db, err := config.GetDB()

    if err != nil {
        log.Fatal(err)
    }

    r := mux.NewRouter()
    uc := controllers.NewUserController(db)

    // Set route pengguna
    routes.SetUserRoutes(r, uc)

    http.Handle("/", r)

    log.Println("Server started on :8080")
    err = http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}
