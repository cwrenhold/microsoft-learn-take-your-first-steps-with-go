package main

import "fmt"

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

func main() {
	employee := Employee{LastName: "Doe", FirstName: "John"}
	fmt.Println(employee)

	// Create a pointer to the employee struct
	employeePointer := &employee
	employeePointer.FirstName = "David"

	// Note: The original employee struct is updated
	fmt.Println(employee)
}
