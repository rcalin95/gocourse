package basics

import "fmt"


func main() {

	// func <name>(<parameters>) <return_type> {
	// 	// function body
	// 	return <value>
	// }

	// sum := add(5, 10)
	// println("Sum is:", sum)

	// greet := func() {
	// 	fmt.Println("Hello Anonymous function!")
	// }

	// greet()

	// operation := add

	// result := operation(20, 30)
	// fmt.Println("Result is:", result)

	// Passing function as parameter
	result := applyOperation(5, 3, add)
	fmt.Println("Applied Operation Result:", result)

	// Reurnong and using function
	multiplyby2 := createMultiplier(2)
	fmt.Println("6 *2:", multiplyby2(6))
}

	func add(a, b int) int {
		return a + b
	}

	//Function that takes another function as parameter
	func applyOperation(a, b int, operation func(int, int) int) int {
		return operation(a, b)
	}

	// Function that retruns another function (the function returned is a closure)
	func createMultiplier(factor int) func(int) int {
		return func(x int) int {
			return x * factor
		}
	}

