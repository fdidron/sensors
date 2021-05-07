package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {
	db, _ = sql.Open("sqlite3", "./db/sensors.db")
	db.Exec("PRAGMA journal_mode=WAL;")

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/sensor/{name}", fetchSensor).Methods("GET", "OPTIONS")
	router.HandleFunc("/sensors/", fetchSensors).Methods("GET", "OPTIONS")
	router.HandleFunc("/reading", createReading).Methods(http.MethodPost)
	fs := http.FileServer(http.Dir("./ui/dist"))
	router.PathPrefix("/").Handler(fs)
	router.Use(mux.CORSMethodMiddleware(router))

	log.Println("Listening on port 8080")
	http.ListenAndServe(":8080", router)
}
