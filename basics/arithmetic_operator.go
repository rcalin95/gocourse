package basics

import (
	"fmt"
	"math"
)

func main() {
	var a, b int = 10, 20
	var result int
	result = a + b
	fmt.Println("Addition:", result)
	result = a - b
	fmt.Println("Subtraction:", result)

	result = a * b
	fmt.Println("Multiplication:", result)

	result = a / b
	fmt.Println("Division:", result)

	result = a % b
	fmt.Println("Modulus:", result)

	const p float64 = 22 / 7.0
	fmt.Println("Constant Pi Approximation:", p)

	// Overflow with signed integers
	var maxInt int64 = 9223337203434243372
	fmt.Println(maxInt)

	maxInt = maxInt + 1
	fmt.Println(maxInt)

	// Overflow with unsigned integers

	var uMaxInt uint64 = 18446744073709551615
	fmt.Println(uMaxInt)

	uMaxInt = uMaxInt + 1
	fmt.Println(uMaxInt)

	var smallFloat float64 = 1.0e-323

	smallFloat = smallFloat / math.MaxFloat64
	fmt.Println("Underflowed Float:", smallFloat)

	

}
