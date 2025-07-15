package backend

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	// Import the database package
	"kai_balai_yasa/backend/database"
)

// Overhaul represents a record in the overhaul table
type Overhaul struct {
	OverhaulID int `json:"overhaul_id"`
	Name string `json:"name"`
	Location string `json:"location"`
	Status string `json:"status"`
	Estimate string `json:"estimate"`
	Progress int `json:"progress"`
	PersonaliaID int `json:"personalia_id"`
	MaterialsID int `json:"materials_id"`
	HistoryID int `json:"history_id"`
	InventoryID int `json:"inventory_id"`
}

// GetOverhaulHandler handles requests to get all overhaul records
func GetOverhaulHandler(w http.ResponseWriter, r *http.Request) {
	db, err := database.InitDB()
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT overhaul_id, name, location, status, estimate, progress, personalia_id, materials_id, history_id, inventory_id FROM overhaul")
	if err != nil {
		log.Printf("Error querying database: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var overhauls []Overhaul
	for rows.Next() {
		var o Overhaul
		if err := rows.Scan(&o.OverhaulID, &o.Name, &o.Location, &o.Status, &o.Estimate, &o.Progress, &o.PersonaliaID, &o.MaterialsID, &o.HistoryID, &o.InventoryID); err != nil {
			log.Printf("Error scanning row: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		overhauls = append(overhauls, o)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(overhauls); err != nil {
		log.Printf("Error encoding JSON: %v", err)
	}
}
