package main

import (
	"fmt"
	"main/mydict"
)

func main() {
	/*	account := accounts.NewAccount("KIM")
		account.Deposit(10)
		fmt.Println(account.Balance())
		err := account.Withdraw(50)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(account.Balance())
		fmt.Println(account)*/
	dictionary := mydict.Dictionary{"first": "First word"}
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
	err = dictionary.Add("second", "Second word")
}
