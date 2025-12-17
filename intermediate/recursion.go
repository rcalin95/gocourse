package intermediate

import "fmt"

func main() {

	fmt.Println(factorial(5))
	fmt.Println(factorial(10))

	fmt.Println(sumOfDigits(9))
	fmt.Println(sumOfDigits(12))
	fmt.Println(sumOfDigits(12344554))

	for i := 0; i < 10; i++ {
		fmt.Print(fibo(i), " ")
	}
}

func factorial(n int) int {
	// Base case : factorial of 0 is 1
	if n == 0 {
		return 1
	}
	// Recursive case : factorial of n is n * factorial(n-1)
	return n * factorial(n-1)
	// n * (n - 1) * (n - 2) * (n - 3) * factorial(n - 4)* ... factorial(0)
}

func sumOfDigits(n int) int {
	// Base case
	if n < 10 {
		return n
	}
	// Recursive case
	return n%10 + sumOfDigits(n/10)
}

func fibo(n int) int {
	// base cases
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	// recursive case
	return fibo(n-1) + fibo(n-2)
	// return nth element of fibonacci series
}
