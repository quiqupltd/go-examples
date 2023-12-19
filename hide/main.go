package main

import (
	"encoding/json"
	"log"

	"github.com/emvi/hide"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Customer struct {
	Id   hide.ID `json:"id" gorm:"primaryKey"`
	Name string  `json:"name" gorm:"not null"`
	Age  int     `json:"age" gorm:"not null"`
}

func main() {
	// Connect to the database
	db, err := connectToSQLite()
	if err != nil {
		panic(err)
	}

	// Migrate the schema
	if err = db.AutoMigrate(&Customer{}); err != nil {
		panic(err)
	}

	// Truuncate table
	if result := db.Exec("DELETE FROM customers"); result.Error != nil {
		panic(result.Error)
	}

	// Create a new customer
	customer := Customer{123, "Foobar", 36}
	if result := db.Create(&customer); result.Error != nil {
		panic(result.Error)
	}

	// Read customer from database
	customers := []Customer{}
	result := db.Find(&customers, customer.Id)

	if result.Error != nil {
		panic(result.Error)
	}

	foundCustomer := customers[0]

	// marshal to JSON
	resultJson, err := json.Marshal(&foundCustomer)
	if err != nil {
		panic(err)
	}

	log.Println(string(resultJson))

	fromJsonCustomer := Customer{}
	// unmarshal from JSON
	if err := json.Unmarshal(resultJson, &fromJsonCustomer); err != nil {
		panic(err)
	}

	log.Println(fromJsonCustomer)
}

func connectToSQLite() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
