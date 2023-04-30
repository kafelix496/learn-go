package main

import (
	"fmt"

	"github.com/kafelix496/learn-go/accounts"
)

func main() {
	account := accounts.NewAccount("kafelix496")
	account.Deposit(10)
	account.Withdraw(3)
	err := account.Withdraw(10)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(account)

	account.ChangeOwner("kafelix28")

	fmt.Println(account)
}
