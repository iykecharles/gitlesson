package db

import (
	"log"
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type ProductItem struct {
	RefPointer int    `sql:"-"`
	tableName  string `sql:"product_items_collections"`
	ID         int    `sql:"id,pk"`
	Name       string `sql:"desc"`
	Desc       string `sql:"desc"`
	Image      string `sql:"image"`
	Features   struct {
		Name string
		Desc string
		Imp  int
	}
	CreatedAt time.Time `sql:"created_at"`
	UpdatedAt time.Time `sql:"updated_at"`
	IsActive  bool      `sql:"is_active"`
}

// CreateProdItemsTable is good to go
func CreateProdItemsTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createErr := db.CreateTable(&ProductItem{}, opts)
	if createErr != nil {
		log.Printf("Table WAS NOT created succesfully, Reason: %v\n", createErr)
		return createErr
	}
	log.Printf("Table HAS BEEN created succesfully.\n")
	return nil

}
