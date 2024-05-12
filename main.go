package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strings"
)

type Address struct {
	Id     int    `json:"id"`
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
	Zip    string `json:"zip"`
}

var (
	stateMap = map[string]string{
		"Alabama": "AL", "Alaska": "AK", "Arizona": "AZ", "Arkansas": "AR", "California": "CA",
		"Colorado": "CO", "Connecticut": "CT", "Delaware": "DE", "Florida": "FL", "Georgia": "GA",
		"Hawaii": "HI", "Idaho": "ID", "Illinois": "IL", "Indiana": "IN", "Iowa": "IA",
		"Kansas": "KS", "Kentucky": "KY", "Louisiana": "LA", "Maine": "ME", "Maryland": "MD",
		"Massachusetts": "MA", "Michigan": "MI", "Minnesota": "MN", "Mississippi": "MS", "Missouri": "MO",
		"Montana": "MT", "Nebraska": "NE", "Nevada": "NV", "New Hampshire": "NH", "New Jersey": "NJ",
		"New Mexico": "NM", "New York": "NY", "North Carolina": "NC", "North Dakota": "ND", "Ohio": "OH",
		"Oklahoma": "OK", "Oregon": "OR", "Pennsylvania": "PA", "Rhode Island": "RI", "South Carolina": "SC",
		"South Dakota": "SD", "Tennessee": "TN", "Texas": "TX", "Utah": "UT", "Vermont": "VT",
		"Virginia": "VA", "Washington": "WA", "West Virginia": "WV", "Wisconsin": "WI", "Wyoming": "WY", "District of Columbia": "DC",
	}
)

func main() {
	db := getDB()
	defer db.Close()

	file, err := os.ReadFile("address.json")
	if err != nil {
		log.Fatal(err)
	}

	var addresses []Address
	err = json.Unmarshal(file, &addresses)
	if err != nil {
		log.Fatal(err)
	}

	err = insertAddresses(db, addresses)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Addresses inserted successfully")
}

func insertAddresses(db *sql.DB, addresses []Address) error {
	var values []string
	for _, addr := range addresses {
		values = append(values, fmt.Sprintf("(%d, '%s', '%s', '%s', '%s')", addr.Id, addr.Street, addr.City, stateMap[addr.State], addr.Zip))
	}

	query := "INSERT INTO address (id, street, city, state, zip) VALUES " + strings.Join(values, ",")
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func getDB() *sql.DB {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=SYS password=password database=SYS sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}
