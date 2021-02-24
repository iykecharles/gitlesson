package db

import (
	"log"
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// ProductItem is good to go
type ProductItem struct {
	RefPointer int    `sql:"-"`
	tableName  string `sql:"product_items_collections"`
	ID         int    `sql:"id,pk"`
	Name       string `sql:"name"`
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

// Save is cool to go
func (pi *ProductItem) Save(db *pg.DB) error {
	insertErr := db.Insert(pi)
	if insertErr != nil {
		log.Printf("Error encountered while inserting data into database. Reason: %v\n", insertErr)
		return insertErr
	}
	log.Printf("Product items have been successfully been inserted into database")
	return nil

}

// GetID is cool to go
func (pi *ProductItem) GetID(db *pg.DB) error {
	getidErr := db.Select(pi)
	if getidErr != nil {
		log.Printf("Error encountered while getting an item from database. Reason: %v\n", getidErr)
		return getidErr
	}
	log.Printf("Product items with stipulated id has been gotten from database. Reason: %v\n", *pi)
	return nil

}

// Update is good to go
func (pi *ProductItem) Update(db *pg.DB) error {
	_, updateErr := db.Model(pi).Set("name = ?name").Where("id = ?id").Update()
	if updateErr != nil {
		log.Printf("Error encountered while updating the name of an item in the database. Reason: %v\n", updateErr)
		return updateErr
	}
	log.Printf("Name of Product items  has been updated in the database. Reason: %v\n", *pi)
	return nil

}

// Delete is good to go
func (pi *ProductItem) Delete(db *pg.DB) error {
	_, deleteErr := db.Model(pi).Where("image = ?image").Delete()
	if deleteErr != nil {
		log.Printf("Error encountered while deleting an item in the database. Reason: %v\n", deleteErr)
		return deleteErr
	}
	log.Printf("Product items  has been deleted in the database for %s", pi.Name)
	return nil

}
