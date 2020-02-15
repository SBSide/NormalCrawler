package main

import (
	"fmt"
	"main/Core"
)

func main() {
	name := 5861
	total := Core.SuperAdd(1, 2, 3, 4, 5, 6)
	totalLength, upperName := Core.LenAndUpper("nyan")
	ignore, _ := Core.LenAndUpper("stars")
	fmt.Println(name)
	fmt.Println(Core.Multiply(5, 25))
	fmt.Println(totalLength, upperName)
	fmt.Println(ignore)
	fmt.Println("SuperAdd : ", total)
}
