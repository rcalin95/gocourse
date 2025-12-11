package basics

import (
	"fmt"
	"slices"
)

func main() {
	// var sliceName []int

	// var numbers = []int
	// var numbers1 = []int{1, 2, 3, 4}

	// numbers2 := []int{5, 6, 7, 8}

	// slice := make([]int, 5)

	a := []int{10, 20, 30, 40, 50}
	slice1 := a[1:4]

	fmt.Println("Slice:", slice1)

	slice1 = append(slice1, 6, 7)
	fmt.Println("Appended Slice:", slice1)

	sliceCopy := make([]int, len(slice1))
	copy(sliceCopy, slice1)

	fmt.Println("Copied Slice:", sliceCopy)

	// var nilSlice []int

	for i,v := range slice1 {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
}

	fmt.Println("Element at index 2:", slice1[2])
	slice1[2] = 100

	fmt.Println("Upldated element at index 2:", slice1[2])

	if slices.Equal(slice1, sliceCopy){
		fmt.Println ("")
	}

	twoD := make ([][]int, 3)
	for i :=0; i<3;i++{
		innerLen := i+1
		twoD[i] = make([]int, innerLen)
		for j :=0;j<innerLen;j++{
			twoD[i][j] = i+j
			fmt.Printf("Adding value %d in outer index %d and inner index %d\n", twoD[i][j], i, j)
		}
	}
	fmt.Println(twoD)

	// slice(low:high)
	slice2 := slice1[2:4] // 4th index is excluded
	fmt.Println(slice2)

	fmt.Println("Capacity of slice2 is:", cap(slice2))
	fmt.Println("Length of slice2 is:", len(slice2))
}