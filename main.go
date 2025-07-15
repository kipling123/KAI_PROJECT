package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"kai_balai_yasa/backend" // Assuming your backend package is named 'backend' within kai_balai_yasa
)

func main() {
	// Initialize the database connection (call this once at the start of your application)
	// Remember to replace the placeholder credentials in backend/database.go
	// db, err := backend.InitDB()
	// if err != nil {
	// 	log.Fatal("Failed to initialize database:", err)
	// }
	// defer db.Close() // Close the database connection when the application exits

	// CORS middleware setup (Allowing all origins for simplicity)
	c := cors.New(cors.Options{ AllowAll: true })

	mux := http.NewServeMux()

	apiHandler := func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		switch r.Method {
		case http.MethodGET:
			switch path {
			case "/inventory/":
				// Handle GET all inventory
				backend.GetInventoryHandler(w, r)
			case "/personalia/":
				backend.GetPersonaliaHandler(w, r)
			case "/produksi/":
				backend.GetProduksiHandler(w, r)
			case "/overhaul/":
				backend.GetOverhaulHandler(w, r)
			case "/rekayasa/":
				backend.GetRekayasaHandler(w, r)
			case "/qualitycontrol/":
				backend.GetQualityControlHandler(w, r)
			case "/stockproduction/":
				backend.GetStockProductionHandler(w, r)
			case "/kalibrasi/":
				backend.GetKalibrasiHandler(w, r)
			case "/profile/":
				backend.GetProfileHandler(w, r)
			default:
				http.NotFound(w, r)
			}
		case http.MethodPOST:
			switch path {
			case "/inventory/":
				backend.CreateInventoryHandler(w, r)
			// Add cases for other POST handlers here
			default:
				http.NotFound(w, r)
			}
		case http.MethodPUT:
			// Logic to parse ID from path and call specific update handlers
			// Example for inventory:
			if strings.HasPrefix(path, "/inventory/") && path != "/inventory/" {
				idStr := strings.TrimPrefix(path, "/inventory/")
				id, err := strconv.Atoi(idStr)
				if err != nil {
					http.Error(w, "Invalid inventory ID", http.StatusBadRequest)
					return
				}
				backend.UpdateInventoryHandler(w, r, id) // Assuming handler accepts ID
			} else {
				// Add cases for other PUT handlers here
				http.NotFound(w, r)

			}
		case http.MethodPOST:
			// Implement POST handlers for each path here or in a nested switch
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}

	// Register the main API handler for the root path.
	// Using HandleFunc with "/" will match all paths that don't have a more specific handler.
	mux.HandleFunc("/", apiHandler)


	// Wrap your main handler with the CORS middleware
	handler := c.Handler(mux)

	log.Println("Server starting on :8080")
	// Use the wrapped handler in ListenAndServe
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}