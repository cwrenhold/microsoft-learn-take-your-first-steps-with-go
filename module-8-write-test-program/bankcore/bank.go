package bank

import (
	"errors"
	"fmt"
)

type Bank interface {
	Statement() string
}

// Customer ...
type Customer struct {
	Name    string
	Address string
	Phone   string
}

// Account ...
type Account struct {
	Customer
	Number  int32
	Balance float64
}

// Deposit ...
func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to deposit should be greater than zero")
	}

	a.Balance += amount
	return nil
}

// Withdraw ...
func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("the amount to withdraw should be greater than zero")
	}

	if a.Balance < amount {
		return errors.New("the amount to withdraw should be greater than the account's balance")
	}

	a.Balance -= amount
	return nil
}

// Statement ...
func (a *Account) Statement() string {
	return fmt.Sprintf("%v - %v - %v", a.Number, a.Name, a.Balance)
}

func Statement(b Bank) string {
	return b.Statement()
}

// Transfer ...
func (a *Account) Transfer(dest *Account, amount float64) error {
	// Withdraw the amount
	if err := a.Withdraw(amount); err != nil {
		return err
	}

	// Deposit the amount
	if err := dest.Deposit(amount); err != nil {
		// In case of error, cancel the withdraw
		a.Deposit(amount)
		return err
	}

	return nil
}
