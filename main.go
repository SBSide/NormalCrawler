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
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(word)
}
