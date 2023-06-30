package onlinestorelogic

import "fmt"

type Employee struct {
	Account
	credits float64
}

func (e *Employee) AddCredits(amount float64) float64 {
	e.credits += amount
	return e.credits
}

func (e *Employee) RemoveCredits(amount float64) float64 {
	e.credits -= amount
	return e.credits
}

func (e *Employee) CheckCredits() float64 {
	return e.credits
}

func (e Employee) String() string {
	return fmt.Sprintf("%s has %.2f credits", e.Account, e.credits)
}
