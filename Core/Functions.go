package Core

import (
	"fmt"
	"strings"
)

func CanIDrink(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false

	case 18:
		return true
	}
	return false
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
