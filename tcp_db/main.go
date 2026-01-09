package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

type Reading struct {
	ID        int    `json:"id,omitempty"`
	MeterID   string `json:"meter_id"`
	Value     int    `json:"value"`
	CreatedAt string `json:"created_at,omitempty"`
}

func main() {

	conn := "host=localhost port=5432 user=postgres password=Sani@123 dbname=tcp_db sslmode=disable"
	if v := os.Getenv("DB_CONN"); v != "" {
		conn = v
	}

	var err error
	db, err = sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Database not connected:", err)
	}

	log.Println("PostgreSQL Connected")
	http.HandleFunc("/save", saveReading)
	http.HandleFunc("/readings", getReadings)

	log.Println("DB Service running on :8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func saveReading(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST only", http.StatusMethodNotAllowed)
		return
	}

	var reading Reading
	if err := json.NewDecoder(r.Body).Decode(&reading); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if reading.MeterID == "" || reading.Value == 0 {
		http.Error(w, "meter_id and value required", http.StatusBadRequest)
		return
	}

	_, err := db.Exec(
		"INSERT INTO meter_readings(meter_id,value) VALUES($1,$2)",
		reading.MeterID, reading.Value,
	)

	if err != nil {
		http.Error(w, "DB Insert Failed", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "saved",
	})
}

func getReadings(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(
		"SELECT id,meter_id,value,created_at FROM meter_readings ORDER BY id DESC LIMIT 50",
	)

	if err != nil {
		http.Error(w, "DB Error", 500)
		return
	}
	defer rows.Close()

	var data []Reading
	for rows.Next() {
		var r Reading
		rows.Scan(&r.ID, &r.MeterID, &r.Value, &r.CreatedAt)
		data = append(data, r)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
