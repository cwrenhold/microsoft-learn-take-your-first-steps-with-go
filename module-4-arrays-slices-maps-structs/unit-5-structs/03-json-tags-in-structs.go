package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	ID int
	// Note: We can use json tags to change the field names
	FirstName string `json:"name"`
	LastName  string
	// Note: We can use json tags to omit fields if empty
	Address string `json:"address,omitempty"`
}

type Employee struct {
	Person
	ManagerID int
}

type Contractor struct {
	Person
	CompanyID int
}

func main() {
	employees := []Employee{
		Employee{
			Person: Person{
				LastName: "Doe", FirstName: "John", Address: "1st Street",
			},
		},
		Employee{
			Person: Person{
				LastName: "Campbell", FirstName: "David",
			},
		},
	}

	// Convert the employees slice to JSON and discard any errors
	data, _ := json.Marshal(employees)
	fmt.Printf("%s\n", data)

	// Convert the JSON data to a slice of Employee structs, ignoring errors
	var decoded []Employee
	json.Unmarshal(data, &decoded)
	fmt.Printf("%v", decoded)
}
