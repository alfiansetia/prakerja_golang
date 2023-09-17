package controllers

import (
    "database/sql"
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "unjuk_keterampilan/models"
)

type UserController struct {
    DB *sql.DB
}

func NewUserController(db *sql.DB) *UserController {
    return &UserController{DB: db}
}

func (uc *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
    users := []models.User{}

    rows, err := uc.DB.Query("SELECT id, name, email FROM users")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    for rows.Next() {
        var user models.User
        err := rows.Scan(&user.ID, &user.Name, &user.Email)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        users = append(users, user)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

func (uc *UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userID := vars["id"]

    user := models.User{}

    row := uc.DB.QueryRow("SELECT id, name, email FROM users WHERE id = ?", userID)
    err := row.Scan(&user.ID, &user.Name, &user.Email)
    if err != nil {
        if err == sql.ErrNoRows {
            // Data tidak ditemukan, kirim respons 404
            w.WriteHeader(http.StatusNotFound)
            return
        }
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}
