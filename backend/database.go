package backend

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// InitDB connects to a MySQL database using the provided credentials.
func InitDB() (*sql.DB, error) {
	// Create the database connection string
	dataSourceName := "your_database_username:your_database_password@tcp(your_database_hostname:3306)/kai_balai_yasa"

	// Open a database connection
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %w", err)
	}

	// Verify the connection is alive
	if err = db.Ping(); err != nil {
		db.Close() // Close the connection if the ping fails
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	fmt.Println("Database connection established successfully!")
	return db, nil
}