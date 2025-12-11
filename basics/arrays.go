package basics

import "fmt"

func main() {

	// var arrayName [size] element type

	// var numbers [5]int
	// fmt.Println(numbers)

	// numbers[4] = 20
	// fmt.Println(numbers)

	// numbers[4] = 9
	// fmt.Println(numbers)

	// fruits := [4]string{"Apple","Banana","Mango","Orange"} 
	// fmt.Println("Fruits array:", fruits)

	// fmt.Println("Third element: ", fruits[2])

	// originalArray := [3]int{1,2,3}
	// copiedArray := originalArray

	// copiedArray[0] = 100

	// fmt.Println("Original Array:", originalArray)
	// fmt.Println("Copied Array:", copiedArray)

	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Printf("Number %d: %d\n", i, numbers[i])
	// }

	// for _, value := range numbers {
	// 	fmt.Printf("Value: %d\n", value)
	// }

	// a, b := someFunction()
	// fmt.Println("Values from someFunction:", a, b)

	// a, _ = someFunction()
	// fmt.Println("Value of a:", a)

	// c := 2
	// _ = c

	// fmt.Println("The length of numbers array is:", len(numbers))

	// //Comparing arrays
	// array1 := [3]int{1,2,3}
	// array2 := [3]int{1,2,3}

	// fmt.Println("Array1 is equal to array2", array1 == array2)

	// var matrix [3][3]int = [3][3]int{
	// 	{1,2,3},
	// 	{4,5,6},
	// 	{7,8,9},
	// }

	// fmt.Println(matrix)

	originalArray := [3]int{1,2,3}
	var copiedArray *[3]int

	copiedArray = &originalArray
	
	copiedArray[0] = 100

	fmt.Println("Original Array:", originalArray)
	fmt.Println("Copied Array:", *copiedArray)

}

func someFunction() (int, int) {
	return 1, 2
}