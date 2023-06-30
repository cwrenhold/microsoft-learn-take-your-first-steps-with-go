package onlinestorelogic

import "fmt"

type Account struct {
	FirstName string
	LastName  string
}

func (a *Account) ChangeName(newFirstName string, newLastName string) {
	a.FirstName = newFirstName
	a.LastName = newLastName
}

func (a Account) String() string {
	return fmt.Sprintf("%s %s", a.FirstName, a.LastName)
}
