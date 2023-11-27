package main

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/mooha76/Puy_me_kofe_API/boot"
	"github.com/mooha76/Puy_me_kofe_API/db"
	mylog "github.com/mooha76/Puy_me_kofe_API/log"
	connection "github.com/mooha76/Puy_me_kofe_API/mongo"
)

func main() {

	cfg, err := connection.Mongoconnection()
	if err != nil {
		log.Fatalf("failed to load config error:%s", err.Error())
	}

	// Call NewConnection from the connection package to get the MongoDB configurations

	// Initialize SQLX Database using the MongoDB configuration

	mylog.InitializeLogger(&cfg.Syslogger)
	sqlxDB, err := db.InitializeSQLXDatabase(&cfg.DBConfig)
	if err != nil {
		log.Fatalf("Error initializing SQLX Database: %v", err)
	}
	defer sqlxDB.Close() // Close the database connection when done

	// Now you can use the sqlxDB object for your database operations
	// For example, you can execute queries, insert data, etc.

	fmt.Println("Successfully initialized SQLX Database!")

	// Print the returned values
	// Print the values
	fmt.Printf("Configuration: %+v\n", cfg)
	fmt.Printf("Error: %v\n", err)
	boot.Boot(&cfg.Sysytemconf)

}
