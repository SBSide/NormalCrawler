package main

import (
	"fmt"
	"main/Core"
)

func main() {
	name := 5861
	totalLength, upperName := Core.LenAndUpper("nyan")
	ignore, _ := Core.LenAndUpper("stars")
	fmt.Println(name)
	fmt.Println(Core.Multiply(5, 25))
	fmt.Println(totalLength, upperName)
	fmt.Println(ignore)
	fmt.Println(Core.CanIDrink(18))
	a := 2
	b := &a
	a = 10
	*b = 20
	fmt.Println(a, *b)

}
