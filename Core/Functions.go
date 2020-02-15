package Core

import (
	"fmt"
	"strings"
)

func SuperAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func Multiply(a, b int) int {
	return a * b
}

func LenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("LenAndUpper : DONE")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}
