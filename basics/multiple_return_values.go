package basics

import (
	"errors"
	"fmt"
)

func main() {
	// func functionName(param1 type1, param2 type2) (returnType1, returnType2) {
	// 	// function body
	// 	return value1, value2
	// }

	q, r := divide(10, 3)
	fmt.Println("Quotient:", q, "Remainder:", r)

	result, err := compare(3, 2)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
}

func divide(a, b int) (quotient int, remainder int) {
	quotient = a / b
	remainder = a % b
	return
}

func compare(a, b int) (string, error) {
	if a > b {
		return "a is greater", nil
	} else if a < b {
		return "b is greater", nil
	} else {
		return "", errors.New("Unable to compare which is greater, both are equal")
	}
}