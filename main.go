package main

import (
	"database/sql"
	_ "embed"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"math/rand"
	"strings"
)

//go:embed employees.json
var employeesContent []byte

//go:embed address.json
var addressContent []byte

//go:embed departments.json
var departmentsContent []byte

var (
	employees   []Employee
	addresses   []Address
	departments []Department
	//colleges []College
)

type Address struct {
	Id     int    `json:"id"`
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
	Zip    string `json:"zip"`
}

type Employee struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	AddressId int    `json:"address_id"`
}

type Department struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	DeanId int    `json:"dean_id"`
}

func init() {
	err := json.Unmarshal(employeesContent, &employees)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(addressContent, &addresses)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(departmentsContent, &departments)
	if err != nil {
		log.Fatal(err)
	}
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
	if false {
		err := insertAddresses(db, addresses)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Addresses inserted successfully")
		err = insertEmployees(db, addresses, employees)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Employees inserted successfully")
	}

	err := insertDepartments(db, employees, departments)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Departments inserted successfully")

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

func insertEmployees(db *sql.DB, addresses []Address, employees []Employee) error {
	var values []string
	for i, addr := range addresses {
		employees[i].AddressId = addr.Id
	}

	args := []any{}

	for i, emp := range employees {
		values = append(values, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d)", (i*6)+1, (i*6)+2, (i*6)+3, (i*6)+4, (i*6)+5, (i*6)+6))
		args = append(args, emp.Id, emp.Name, emp.Phone, emp.Email, emp.Gender, emp.AddressId)
	}

	query := "INSERT INTO employee (emp_id, name, phone, email, gender, address_id) VALUES " + strings.Join(values, ",")
	_, err := db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func insertDepartments(db *sql.DB, employees []Employee, departments []Department) error {
	var values []string

	args := []any{}
	for i, dep := range departments {
		randomIndex := rand.Intn(1001)
		id := employees[randomIndex].Id
		departments[i].DeanId = id
		values = append(values, fmt.Sprintf("($%d, $%d, $%d)", (i*3)+1, (i*3)+2, (i*3)+3))
		args = append(args, dep.Id, dep.Name, id)
	}

	query := "INSERT INTO department (id, name, dean_id) VALUES " + strings.Join(values, ",")
	_, err := db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil

}

/*func insertColleges(db *sql.DB, colleges []College) error {
	var values []string

	args := []any{}

}*/

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
