package backend

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql" // Import the MySQL driver
)

type Personalia struct {
	PersonaliaID int    `json:"personalia_id"`
	NIP          int    `json:"nip"`
	Jabatan      string `json:"jabatan"`
	Divisi       string `json:"divisi"`
	Lokasi       string `json:"lokasi"`
	Status       string `json:"status"`
	JoinDate     string `json:"join_date"`
	PhoneNumber  string `json:"phone_number"`
	UrgentNumber string `json:"urgent_number"`
	ProfileID    int    `json:"profile_id"`
}

func GetPersonaliaHandler(w http.ResponseWriter, r *http.Request) {
	db, err := InitDB() // Assuming InitDB is in database.go and accessible
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT personalia_id, nip, jabatan, divisi, lokasi, status, join_date, phone_number, urgent_number, profile_id FROM personalia")
	if err != nil {
		log.Printf("Error querying personalia table: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	personaliaList := []Personalia{}
	for rows.Next() {
		var p Personalia
		if err := rows.Scan(&p.PersonaliaID, &p.NIP, &p.Jabatan, &p.Divisi, &p.Lokasi, &p.Status, &p.JoinDate, &p.PhoneNumber, &p.UrgentNumber, &p.ProfileID); err != nil {
			log.Printf("Error scanning personalia row: %v", err)
			continue // Skip this row, log the error
		}
		personaliaList = append(personaliaList, p)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(personaliaList); err != nil {
		log.Printf("Error encoding personalia list: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

// GetPersonaliaByIDHandler handles requests to get a single personalia entry by ID.
func GetPersonaliaByIDHandler(w http.ResponseWriter, r *http.Request, personaliaID int) {
	db, err := InitDB() // Assuming InitDB is in database.go and accessible
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var p Personalia
	// Query for a single personalia entry by ID
	row := db.QueryRow("SELECT personalia_id, nip, jabatan, divisi, lokasi, status, join_date, phone_number, urgent_number, profile_id FROM personalia WHERE personalia_id = ?", personaliaID)

	err = row.Scan(&p.PersonaliaID, &p.NIP, &p.Jabatan, &p.Divisi, &p.Lokasi, &p.Status, &p.JoinDate, &p.PhoneNumber, &p.UrgentNumber, &p.ProfileID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Personalia not found", http.StatusNotFound)
		} else {
			log.Printf("Error scanning personalia row: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(p); err != nil {
		log.Printf("Error encoding personalia: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

// CreatePersonaliaHandler handles requests to create a new personalia entry.
//
// CreatePersonaliaHandler handles requests to create a new personalia entry.
func CreatePersonaliaHandler(w http.ResponseWriter, r *http.Request) {
	db, err := InitDB() // Assuming InitDB is in database.go and accessible
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer db.Close()
	
	var newPersonalia Personalia
	// Decode the JSON request body into a Personalia struct.
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newPersonalia)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}


	// TODO: Implement SQL INSERT statement to insert newPersonalia data into the 'personalia' table.
	// Example:
	// stmt, err := db.Prepare("INSERT INTO personalia (nip, jabatan, divisi, lokasi, status, join_date, phone_number, urgent_number, profile_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)")
	// if err != nil {
	// 	log.Printf("Error preparing insert statement: %v", err)
	// 	http.Error(w, "Internal server error", http.StatusInternalServerError)
	// 	return
	// }
	// defer stmt.Close()
	//
	// result, err := stmt.Exec(newPersonalia.NIP, newPersonalia.Jabatan, newPersonalia.Divisi, newPersonalia.Lokasi, newPersonalia.Status, newPersonalia.JoinDate, newPersonalia.PhoneNumber, newPersonalia.UrgentNumber, newPersonalia.ProfileID)
	// if err != nil {
	// 	log.Printf("Error executing insert statement: %v", err)
	// 	http.Error(w, "Internal server error", http.StatusInternalServerError)
	// 	return
	// }

	// SQL INSERT statement
	insertStatement := "INSERT INTO personalia (nip, jabatan, divisi, lokasi, status, join_date, phone_number, urgent_number, profile_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	
	// Execute the INSERT statement
	result, err := db.Exec(insertStatement,
		newPersonalia.NIP,
		newPersonalia.Jabatan,
		newPersonalia.Divisi,
		newPersonalia.Lokasi,
		newPersonalia.Status,
		newPersonalia.JoinDate, // Ensure JoinDate is in a format compatible with your SQL DATE type
		newPersonalia.PhoneNumber,
		newPersonalia.UrgentNumber,
		newPersonalia.ProfileID,
	)
	if err != nil {
		log.Printf("Error executing insert statement: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// TODO: Respond with success, possibly including the ID of the newly created entry.
	// w.WriteHeader(http.StatusCreated)
	// json.NewEncoder(w).Encode(map[string]string{"message": "Personalia created successfully"})
}

// UpdatePersonaliaHandler handles requests to update an existing personalia entry.
// It takes the personalia ID to update as an integer argument.
func UpdatePersonaliaHandler(w http.ResponseWriter, r *http.Request, personaliaID int) {
	db, err := InitDB()
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer db.Close() // Close the database connection when the function finishes

	var updatedPersonalia Personalia // Use a new variable for decoding
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&updatedPersonalia)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest) // Return a 400 for bad request payload
		return
	}

	// SQL UPDATE statement
	// Use the ID from the URL for the WHERE clause, not from the decoded body unless you want to allow updating the ID (usually not recommended)
	// Also, ensure the order of arguments in Exec matches the order of placeholders in the UPDATE statement
	updateStatement := `UPDATE personalia SET nip = ?, jabatan = ?, divisi = ?, lokasi = ?, status = ?, join_date = ?, phone_number = ?, urgent_number = ?, profile_id = ? WHERE personalia_id = ?`

	// Execute the UPDATE statement
	result, err := db.Exec(updateStatement,
		updatedPersonalia.NIP,
		updatedPersonalia.Jabatan,
		updatedPersonalia.Divisi,
		updatedPersonalia.Lokasi,
		updatedPersonalia.Status,
		updatedPersonalia.JoinDate, // Ensure JoinDate is in a format compatible with your SQL DATE type
		updatedPersonalia.PhoneNumber,
		updatedPersonalia.UrgentNumber,
		updatedPersonalia.ProfileID,
		personaliaID, // Use the ID from the URL
	)
	if err != nil {
		log.Printf("Error executing update statement: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError) // Return a 500 for internal server errors
		return
	}

	// Check the number of rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Personalia not found", http.StatusNotFound) // Return 404 if no rows were updated
		return
	}

	// TODO: Respond with success message.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // Return 200 OK for successful update
	json.NewEncoder(w).Encode(map[string]string{"message": "Personalia updated successfully"}) // Respond with a success message
}

// DeletePersonaliaHandler handles requests to delete an existing personalia entry.
// It takes the personalia ID to delete as an integer argument.
func DeletePersonaliaHandler(w http.ResponseWriter, r *http.Request, personaliaID int) {
	db, err := InitDB()
	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer db.Close() // Close the database connection when the function finishes
	//
	// SQL DELETE statement
	deleteStatement := "DELETE FROM personalia WHERE personalia_id = ?"

	// Execute the DELETE statement
	result, err := db.Exec(deleteStatement, personaliaID)
	if err != nil {
		log.Printf("Error executing delete statement: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError) // Return a 500 for internal server errors
		return
	}

	// Check if any rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Personalia not found", http.StatusNotFound) // Return 404 if no rows were deleted
		return
	}

	rowsAffected, err := result.RowsAffected()
	// 	return
	// }

	// TODO: Respond with success message.
	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(map[string]string{"message": "Personalia deleted successfully"})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // Return 200 OK for successful deletion
	json.NewEncoder(w).Encode(map[string]string{"message": "Personalia deleted successfully"}) // Respond with a success message
}

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

type Education struct {
	ID         int    `json:"id"`
	Degree     string `json:"degree"`
	University string `json:"university"`
	Year       string `json:"year"`
}

type Experience struct {
	ID       int    `json:"id"`
	Position string `json:"position"`
	Company  string `json:"company"`
	Period   string `json:"period"`
}

type Profile struct {
	Name       string       `json:"name"`
	Position   string       `json:"position"`
	Email      string       `json:"email"`
	Phone      string       `json:"phone"`
	Address    string       `json:"address"`
	Bio        string       `json:"bio"`
	Education  []Education  `json:"education"`
	Experience []Experience `json:"experience"`
}

var (
	profile = Profile{
		Name:     "Asep Hidayat S.Kom M.Kom",
		Position: "Production Manager",
		Email:    "john.doe@kai.co.id",
		Phone:    "+62 812-3456-7890",
		Address:  "Jl. Kereta Api No. 1, Jakarta",
		Bio:      "Professional with 10+ years experience in railway production management",
		Education: []Education{
			{ID: 1, Degree: "Bachelor of Engineering", University: "Institut Teknologi Bandung", Year: "2005-2009"},
			{ID: 2, Degree: "Master of Business Administration", University: "Universitas Indonesia", Year: "2011-2013"},
		},
		Experience: []Experience{
			{ID: 1, Position: "Production Supervisor", Company: "PT KAI", Period: "2010-2015"},
			{ID: 2, Position: "Production Manager", Company: "PT KAI Balai Yasa", Period: "2015-Present"},
		},
	}
	mutex sync.Mutex
)

func enableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func getProfile(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	if r.Method == http.MethodOptions {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	mutex.Lock()
	defer mutex.Unlock()
	if err := json.NewEncoder(w).Encode(profile); err != nil {
		http.Error(w, "Failed to encode profile", http.StatusInternalServerError)
	}
}

func updateProfile(w http.ResponseWriter, r *http.Request) {
	enableCORS(w)
	if r.Method == http.MethodOptions {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	mutex.Lock()
	defer mutex.Unlock()

	var updatedProfile Profile
	if err := json.NewDecoder(r.Body).Decode(&updatedProfile); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	profile = updatedProfile

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(profile); err != nil {
		http.Error(w, "Failed to encode updated profile", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getProfile(w, r)
		case http.MethodPut:
			updateProfile(w, r)
		case http.MethodOptions:
			enableCORS(w)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Personalia API running on http://localhost:8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
