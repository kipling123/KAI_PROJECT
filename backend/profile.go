package backend

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"kai_balai_yasa/backend/database" // Sesuaikan dengan path import database Anda
)

// Profile represents a row in the 'profile' table.
type Profile struct {
	ProfileID int    `json:"profile_id"`
	Email     string `json:"email"`
	Address   string `json:"addres"` // Note the spelling based on the schema
	PhoneNumber string `json:"phone_number"`
	EducationID int    `json:"education_id"`
	ExperienceID int    `json:"experience_id"`
}

// GetProfileHandler handles requests to get all profile data.
func GetProfileHandler(w http.ResponseWriter, r *http.Request) {
	db, err := database.InitDB()
	if err != nil {
		log.Printf("Error initializing database: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT profile_id, email, addres, phone_number, education_id, experience_id FROM profile")
	if err != nil {
		log.Printf("Error querying profile table: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	profiles := []Profile{}
	for rows.Next() {
		var p Profile
		if err := rows.Scan(&p.ProfileID, &p.Email, &p.Address, &p.PhoneNumber, &p.EducationID, &p.ExperienceID); err != nil {
			log.Printf("Error scanning profile row: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		profiles = append(profiles, p)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error during rows iteration: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(profiles); err != nil {
		log.Printf("Error encoding profiles to JSON: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}