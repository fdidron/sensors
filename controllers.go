package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func fetchSensor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]

	s := Sensor{}
	s.FindWithName(name)
	respondWithJSON(w, http.StatusOK, s)
}

func fetchSensors(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, getSensors())
}

type ReadingPayload struct {
	Sensor string `json:"sensor"`
	Value  string `json:"value"`
}

func createReading(w http.ResponseWriter, r *http.Request) {

	var re ReadingPayload
	d := json.NewDecoder(r.Body)

	if err := d.Decode(&re); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		log.Println(err)
		return
	}

	s := Sensor{}
	s.FindWithName(re.Sensor)

	if s.Id == 0 {
		respondWithError(w, http.StatusNotFound, "Sensor not found")
		return
	}

	reading := Reading{Value: re.Value, SensorID: s.Id}
	reading.Save()

	respondWithJSON(w, http.StatusCreated, reading)
}
