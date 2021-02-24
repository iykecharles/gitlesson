package db

import (
	"log"
	"os"

	"github.com/go-pg/pg"
)

// Connect would connect to the postgres database
func Connect() *pg.DB{
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
	return db

}
