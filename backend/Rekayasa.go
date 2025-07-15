package backend

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Rekayasa struct {
	RekayasaID int    `json:"rekayasa_id"`
	Name       string `json:"name"`
	Status     string `json:"status"`
	Team       string `json:"team"` // Assuming team is stored as a string (e.g., comma-separated)
	Deadline   string `json:"deadline"`
	Progress   string `json:"progress"` // Assuming progress is stored as varchar
}

// GetRekayasaHandler handles requests to get all rekayasa records.
func GetRekayasaHandler(w http.ResponseWriter, r *http.Request) {
	db, err := InitDB()
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT rekayasa_id, name, status, team, deadline, progress FROM rekayasa")
	if err != nil {
		log.Printf("Error querying rekayasa table: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var rekayasaList []Rekayasa
	for rows.Next() {
		var rk Rekayasa
		if err := rows.Scan(&rk.RekayasaID, &rk.Name, &rk.Status, &rk.Team, &rk.Deadline, &rk.Progress); err != nil {
			log.Printf("Error scanning rekayasa row: %v", err)
			continue // Skip this row and try the next
		}
		rekayasaList = append(rekayasaList, rk)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rekayasaList)
}
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"


)

type Project struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Status   string   `json:"status"`
	Team     []string `json:"team"`
	Deadline string   `json:"deadline"`
	Progress int      `json:"progress"`
}

var projects = []Project{
	{
		ID:       1,
		Name:     "Pengembangan Sistem Kontrol",
		Status:   "Dalam Pengerjaan",
		Team:     []string{"BS", "AW", "CD"},
		Deadline: "2023-12-31",
		Progress: 65,
	},
	{
		ID:       2,
		Name:     "Optimasi Produksi",
		Status:   "Selesai",
		Team:     []string{"DP", "ES"},
		Deadline: "2023-10-15",
		Progress: 100,
	},
	{
		ID:       3,
		Name:     "Desain Komponen Baru",
		Status:   "Perencanaan",
		Team:     []string{"BS", "ES"},
		Deadline: "2024-02-28"
		Progress: 15,
	},
}

func getProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}

func main() {
	http.HandleFunc("/api/projects", getProjects)

	log.Println("Server berjalan di http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
