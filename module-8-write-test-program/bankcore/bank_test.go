package bank

import "testing"

func getTestAccount() Account {
	return Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}
}

func TestAccount(t *testing.T) {
	account := getTestAccount()

	if account.Name == "" {
		t.Error("can't create an Account object")
	}
}

func TestDeposit(t *testing.T) {
	account := getTestAccount()

	account.Deposit(10)

	if account.Balance != 10 {
		t.Error("balance is not being updated after a deposit")
	}
}

func TestDepositInvalid(t *testing.T) {
	account := getTestAccount()

	if err := account.Deposit(-10); err == nil {
		t.Error("only positive numbers should be allowed to deposit")
	}
}

func TestWithdraw(t *testing.T) {
	account := getTestAccount()

	account.Deposit(10)
	account.Withdraw(10)

	if account.Balance != 0 {
		t.Error("balance is not being updated after withdraw")
	}
}

func TestWithdrawNegative(t *testing.T) {
	account := getTestAccount()

	account.Deposit(10)

	if err := account.Withdraw(-10); err == nil {
		t.Error("only positive numbers should be allowed to withdraw")
	}
}

func TestWithdrawWithNoBalance(t *testing.T) {
	account := getTestAccount()

	if err := account.Withdraw(10); err == nil {
		t.Error("should not be allowed to withdraw with no balance")
	}
}

func TestStatement(t *testing.T) {
	account := getTestAccount()

	account.Deposit(100)
	statement := account.Statement()
	if statement != "1001 - John - 100" {
		t.Error("statement doesn't have the proper format")
	}
}

func TestTransfer(t *testing.T) {
	account1 := getTestAccount()
	account2 := getTestAccount()
	account2.Number = 1002

	account1.Deposit(100)
	account1.Transfer(&account2, 100)

	if account1.Balance != 0 || account2.Balance != 100 {
		t.Error("transfer is not working")
	}
}
