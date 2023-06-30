package main

import (
	"fmt"

	"github.com/cwrenhold/onlinestorelogic"
)

func main() {
	employee := onlinestorelogic.Employee{
		Account: onlinestorelogic.Account{
			FirstName: "John",
			LastName:  "Doe",
		},
	}

	fmt.Println(employee)
	fmt.Println("Employee credits: ", employee.CheckCredits())

	employee.AddCredits(10)
	fmt.Println("Employee credits: ", employee.CheckCredits())

	employee.RemoveCredits(5)
	fmt.Println("Employee credits: ", employee.CheckCredits())

	employee.ChangeName("Jane", "Doe")
	fmt.Println(employee)
}
