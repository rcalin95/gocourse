package basics

import (
	"fmt"
)

func main() {

	// func funcitonName (param1 type1, param2 type2, param3 ...type) returnType {
	//  function body
	// 	return returnValue
	// }

	// fmt.Println("Sum:", sum(1, 2, 3))
	// statement, total := sum("Total Sum:", 1, 2, 3, 4, 5)
	// fmt.Println(statement, total)
	sequence, total := sum(1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("Sequence:", sequence, "Total:", total)
	sequence2, total2 := sum(2, 10, 20, 30)
	fmt.Println("Sequence:", sequence2, "Total:", total2)

	numbers := []int{1, 2, 3, 4, 5, 9} // slice of integers

	sequence3, total3 := sum(3, numbers...)
	fmt.Println("Sequence:", sequence3, "Total:", total3)
}

func sum(sequence int, nums ...int) (int, int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	return sequence, total
}
