package db

import (
	"log"
	"os"

	"github.com/go-pg/pg"
)

// Connect would connect to the postgres database
func Connect() {
	opts := &pg.Options{
		Addr:     "localhost:5432",
		User:     "postgres",
		Password: "charles",
		Database: "goweb",
	}
	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Database DID NOT connect succesfully")
		os.Exit(100)

	}
	log.Printf("Database connected succesfully")
	
	CreateProdItemsTable(db)
	closeErr := db.Close()
	if closeErr != nil {
		log.Printf("Erroe found while closing database")
		os.Exit(100)
	}
	log.Printf("Database has been closed successfully")
	return

}
