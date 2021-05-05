package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

var db *sql.DB

func main() {
	db, _ = sql.Open("sqlite3", "./db/sensors.db")
	db.Exec("PRAGMA journal_mode=WAL;")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/sensor/{name}", fetchSensor).Methods("GET")
	router.HandleFunc("/sensors/", fetchSensors).Methods("GET")
	router.HandleFunc("/reading", createReading).Methods("POST")
	fs := http.FileServer(http.Dir("./ui/dist"))
	router.PathPrefix("/").Handler(fs)

	log.Println("Listening on port 8080")
	http.ListenAndServe(":8080", router)
}
