package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mooha76/Puy_me_kofe_API/config"
)

func InitializeSQLXDatabase(cfg *config.DatabaseConfig) (*sqlx.DB, error) {
	// Construct the DSN (Data Source Name) for PostgreSQL
	dsn := fmt.Sprintf(
		"host=%s user=%s dbname=%s port=%s password=%s sslmode=%s",
		cfg.DB_HOST, cfg.DB_USER, cfg.DB_NAME, cfg.DB_PORT, cfg.DB_PASSWORD, cfg.Sslmode,
	)

	// Open a connection to the PostgreSQL database
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	// Ping the database to check if the connection is successful
	if err := db.Ping(); err != nil {
		db.Close() // Close the database connection in case of an error
		return nil, fmt.Errorf("error pinging database: %v", err)
	}

	return db, nil
}
