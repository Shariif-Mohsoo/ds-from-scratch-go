package embedded

import "fmt"

type Account struct {
	accId   int
	balance int
	name    string
}

func (a *Account) GetBalance() int {
	return a.balance
}

type ManagerAccount struct {
	Account
}

func Embedded() {
	mgrAccount := ManagerAccount{Account{2, 30, "Cassandra"}}
	fmt.Println(mgrAccount)
	fmt.Println("Balance:", mgrAccount.GetBalance())
}
