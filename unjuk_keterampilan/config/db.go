package config

import (
    "database/sql"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

// GetDB berfungsi untuk menginisialisasi dan mengembalikan koneksi database
func GetDB() (*sql.DB, error) {
    // Ganti informasi koneksi database sesuai dengan konfigurasi Anda
    username := "root"
    password := ""
    databaseName := "golang"
    host := "localhost"
    port := "3306"

    // Buat string koneksi
    dataSourceName := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + databaseName

    // Inisialisasi koneksi database
    db, err := sql.Open("mysql", dataSourceName)
    if err != nil {
        log.Fatal(err)
        return nil, err
    }

    // Tes koneksi ke database
    err = db.Ping()
    if err != nil {
        log.Fatal(err)
        return nil, err
    }

    return db, nil
}
