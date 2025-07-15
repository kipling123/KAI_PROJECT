package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv" // Added import for strconv
	"time"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux" // Assuming you'll use gorilla/mux for routing and path variables

	// Assuming database.go is in the same 'backend' package or imported correctly
	"your_module_path/backend/database" // Replace with your actual module path
)

type Inventory struct {
	InventoryID int64 `json:"inventory_id"` // Changed to int64 for LastInsertId compatibility
	Name        string `json:"name"`
	Location  string    `json:"location"`
	Status    string    `json:"status"`
	Qty         int       `json:"qty"` // Added Qty field based on DB schema
}

// GetInventoryHandler handles GET requests for inventory data.
func GetInventoryHandler(w http.ResponseWriter, r *http.Request) {
	db, err := database.InitDB() // Assuming InitDB is in database.go and returns *sql.DB and error
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT inventory_id, name, qty, location, status FROM inventory") // Added qty to query
	if err != nil { // Removed extra parenthesis
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var inventoryItems []Inventory // Corrected variable name
	for rows.Next() {
		var item Inventory // Declared item inside the loop
		if err := rows.Scan(&item.InventoryID, &item.Name, &item.Qty, &item.Location, &item.Status); err != nil {
			http.Error(w, "Failed to scan row", http.StatusInternalServerError) // Fixed typo
			return
		}
		inventoryItems = append(inventoryItems, item) // Corrected append
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(inventoryItems); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		log.Printf("Error encoding response: %v", err) // Added logging for encoding error
	}
}

// GetInventoryByIDHandler handles GET requests for a single inventory item by ID.
func GetInventoryByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		log.Printf("Error decoding request body: %v", err)
		return
	}

	db, err := InitDB() // Get DB connection
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		log.Printf("Error connecting to database: %v", err)
		return
	}
	defer db.Close()
}

// CreateInventoryHandler handles POST requests to add new inventory items.
func CreateInventoryHandler(w http.ResponseWriter, r *http.Request) {
	var newInventory Inventory // Use the Inventory struct, changed variable name to newInventory
	if err := json.NewDecoder(r.Body).Decode(&newInventory); err != nil { // Decode into newInventory
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		log.Printf("Error decoding request body: %v", err)
		return
	}

	db, err := database.InitDB() // Get DB connection
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		log.Printf("Error connecting to database: %v", err)
		return
	}
	defer db.Close()

	// Prepare the SQL INSERT statement
	insertStatement := "INSERT INTO inventory (name, qty, location, status) VALUES (?, ?, ?, ?)" // Changed to insertStatement
	result, err := db.Exec(insertStatement, newInventory.Name, newInventory.Qty, newInventory.Location, newInventory.Status) // Execute the statement directly
	if err != nil {
		http.Error(w, "Failed to create inventory item", http.StatusInternalServerError)
		log.Printf("Error executing insert statement: %v", err)
		return
	}

	// Get the ID of the newly inserted item (optional but good practice)
	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Warning: Could not get last insert ID: %v", err)
	}
	newInventory.InventoryID = id // Set the ID in the response struct

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // HTTP 201 Created
	if err := json.NewEncoder(w).Encode(newInventory); err != nil { // Encode newInventory
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		log.Printf("Error encoding response: %v", err)
	}
}

// UpdateInventoryHandler handles PUT requests to update existing inventory items.
func UpdateInventoryHandler(w http.ResponseWriter, r *http.Request) {
	// Use mux.Vars for extracting ID from URL path
	// Extract inventory ID from URL path using gorilla/mux
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "Inventory ID not provided in URL", http.StatusBadRequest)
		return // Corrected return position
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Inventory ID", http.StatusBadRequest)
		log.Printf("Error converting ID to int: %v", err)
		return
	}

	var updatedItem Inventory // Use the Inventory struct
	if err := json.NewDecoder(r.Body).Decode(&updatedItem); err != nil { // Decode into updatedItem
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		log.Printf("Error decoding request body: %v", err)
		return
	}

	// Ensure the ID from the URL matches the ID in the request body if present
	if updatedItem.InventoryID != 0 && updatedItem.InventoryID != int64(id) {
		http.Error(w, "Mismatch between URL ID and body ID", http.StatusBadRequest) // Corrected error message
		return
	}
	updatedItem.InventoryID = int64(id) // Use the ID from the URL

	db, err := database.InitDB() // Get DB connection
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		log.Printf("Error connecting to database: %v", err)
		return
	}
	defer db.Close()

	// Prepare the SQL UPDATE statement
	result, err := db.Exec(insertStatement, newInventory.Name, newInventory.Qty, newInventory.Location, newInventory.Status) // Execute the statement directly
		if err != nil {
		http.Error(w, "Failed to create inventory item", http.StatusInternalServerError)
		log.Printf("Error executing insert statement: %v", err)
		return
	}
	
	// Get the ID of the newly inserted item (optional but good practice)
	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Warning: Could not get last insert ID: %v", err)
	}
	newInventory.InventoryID = id // Set the ID in the response struct
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // HTTP 201 Created
	if err := json.NewEncoder(w).Encode(newInventory); err != nil { // Encode newInventory
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		log.Printf("Error encoding response: %v", err)
	}
}

// UpdateInventoryHandler handles PUT requests to update existing inventory items.
func UpdateInventoryHandler(w http.ResponseWriter, r *http.Request) {
	// Use mux.Vars for extracting ID from URL path
	// Extract inventory ID from URL path using gorilla/mux
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "Inventory ID not provided in URL", http.StatusBadRequest)
		return // Corrected return position
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Inventory ID", http.StatusBadRequest)
		log.Printf("Error converting ID to int: %v", err)
		return
	}

	var updatedItem Inventory // Use the Inventory struct
	if err := json.NewDecoder(r.Body).Decode(&updatedItem); err != nil { // Decode into updatedItem
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		log.Printf("Error decoding request body: %v", err)
		return
	}

	// Ensure the ID from the URL matches the ID in the request body if present and not zero
	if updatedItem.InventoryID != 0 && updatedItem.InventoryID != int64(id) {
		http.Error(w, "Mismatch between URL ID and body ID", http.StatusBadRequest) // Corrected error message
		return
	}
	updatedItem.InventoryID = id // Use the ID from the URL

	db, err := InitDB() // Get DB connection
		if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		log.Printf("Error connecting to database: %v", err)
		return
	}
	defer db.Close()
	
	// Prepare the SQL UPDATE statement
	updateStatement := `UPDATE inventory SET name=?, qty=?, location=?, status=? WHERE inventory_id=?` // Use qty, update statement
	_, err = db.Exec(updateStatement, updatedItem.Name, updatedItem.Qty, updatedItem.Location, updatedItem.Status, updatedItem.InventoryID) // Execute with updatedItem data and ID
		if err != nil {
			http.Error(w, "Failed to update inventory item", http.StatusInternalServerError)
			log.Printf("Error executing update statement: %v", err)
			return
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			log.Printf("Warning: Could not get rows affected: %v", err)
		}

		if rowsAffected == 0 {
			http.Error(w, "Inventory item not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) // HTTP 200 OK
		response := map[string]string{"message": "Inventory updated successfully"}
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			log.Printf("Error encoding response: %v", err)
		}
}

// DeleteInventoryHandler handles DELETE requests to delete inventory items.
func DeleteInventoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "Inventory ID not provided in URL", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		log.Printf("Error converting ID to int for delete: %v", err)
		return
	}

	db, err := InitDB() // Get DB connection
	if err != nil {
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		log.Printf("Error connecting to database for delete: %v", err)
		return
	}
}

import "fmt"

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}
