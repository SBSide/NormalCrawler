package main

import (
	"fmt"
	"main/banking"
)

func main() {
	/*	name := 5861
		totalLength, upperName := core.LenAndUpper("nyan")
		ignore, _ := core.LenAndUpper("stars")
		fmt.Println(name)
		fmt.Println(core.Multiply(5, 25))
		fmt.Println(totalLength, upperName)
		fmt.Println(ignore)
		fmt.Println(core.CanIDrink(18))
		a := 2
		b := &a
		a = 10
		*b = 20
		names := []string{"nyan", "nay", "yaong"}
		names = append(names, "meow")
		dogs := map[string]string{"dogs": "bark", "name": "shiro"}
		for key, val := range dogs {
			println(key, val)
		}
		fmt.Println(a, *b)
		fmt.Println(names, dogs)
		favFood := []string{"pizza", "ramen"}
		me := person{
			name:    "KIM",
			age:     26,
			favFood: favFood,
		}
		fmt.Println(me.name)*/
	account := banking.Account{
		Owner:   "Nyan",
		Balance: 10000,
	}
	fmt.Println(account)
}
