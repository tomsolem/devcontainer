package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

type Equipment struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Quantity int    `json:"quantity"`
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("postgres", "user=postgres dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, world!")
	})

	http.HandleFunc("/equipment", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM equipment")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var equipments []Equipment
		for rows.Next() {
			var e Equipment
			if err := rows.Scan(&e.ID, &e.Name, &e.Type, &e.Quantity); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			equipments = append(equipments, e)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(equipments)
	})

	http.ListenAndServe(":5000", nil)
}
