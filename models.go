package main

import (
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Sensor struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	CategoryID  string  `json:"categoryId"`
	LastReading Reading `json:"lastReading"`
}

func (s *Sensor) FindWithName(name string) {
	s.Name = name
	rows, _ := db.Query(`
SELECT s.id, s.name, s.description, s.category_id, r.id, r.value, r.created_at
FROM sensors s LEFT JOIN readings r ON r.sensor_id = s.id
WHERE s.name = ?
ORDER BY r.created_at DESC`, name)
	defer rows.Close()

	rows.Next()
	rows.Scan(&s.Id, &s.Name, &s.Description, &s.CategoryID,
		&s.LastReading.Id, &s.LastReading.Value, &s.LastReading.CreatedAt)
}

func getSensors() []Sensor {
	var sensors []Sensor
	rows, err := db.Query(`
SELECT s.id, s.name, s.description, s.category_id, r.id, r.value, r.created_at
FROM sensors s LEFT JOIN readings r ON r.id = (SELECT id FROM readings WHERE sensor_id = s.id ORDER BY created_at DESC LIMIT 1)`)
	if err != nil {
		log.Println(err)
		return sensors
	}
	defer rows.Close()

	for rows.Next() {
		s := Sensor{}
		rows.Scan(&s.Id, &s.Name, &s.Description, &s.CategoryID,
			&s.LastReading.Id, &s.LastReading.Value, &s.LastReading.CreatedAt)
		sensors = append(sensors, s)
	}

	return sensors
}

type Reading struct {
	Id        int    `json:"id"`
	Value     string `json:"value"`
	SensorID  int    `json:"sensorId"`
	CreatedAt string `json:"createdAt"`
}

func (r *Reading) Save() {
	stmt, _ := db.Prepare("INSERT INTO readings(value, sensor_id) values(?,?)")

	res, err := stmt.Exec(r.Value, r.SensorID)
	if err != nil {
		log.Println(err)
	}
	id, _ := res.LastInsertId()
	r.Id = int(id)
}
