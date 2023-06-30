package main

import "fmt"

type Person struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
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
	// Note: We initialise the Employee struct with the Person struct
	employee := Employee{
		Person: Person{
			FirstName: "John",
		},
	}

	// Note: We can access the Person struct fields directly from the Employee struct
	employee.LastName = "Doe"
	fmt.Println(employee.FirstName)
}
