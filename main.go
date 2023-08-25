package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

type Route struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Source      Location `json:"source"`
	Destination Location `json:"destination"`
}

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "root:password@tcp(db:3306)/routesdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := chi.NewRouter()
	r.Post("/api/routes", createRoute)
	r.Get("/api/routes", listRoutes)

	http.ListenAndServe(":8080", r)
}

func createRoute(w http.ResponseWriter, r *http.Request) {
	var route Route
	err := json.NewDecoder(r.Body).Decode(&route)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Inserir a rota no banco de dados
	_, err = db.Exec("INSERT INTO routes (name, source_lat, source_lng, destination_lat, destination_lng) VALUES (?, ?, ?, ?, ?)",
		route.Name, route.Source.Lat, route.Source.Lng, route.Destination.Lat, route.Destination.Lng)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func listRoutes(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, source_lat, source_lng, destination_lat, destination_lng FROM routes")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var routes []Route
	for rows.Next() {
		var route Route
		err := rows.Scan(&route.ID, &route.Name, &route.Source.Lat, &route.Source.Lng, &route.Destination.Lat, &route.Destination.Lng)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		routes = append(routes, route)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(routes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
