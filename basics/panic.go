package basics

import "fmt"

func main() {

	// panic(interface{})
	process(-3)
}

func process(i int) {

	defer fmt.Println("Deffered 1")
	defer fmt.Println("Deffered 2")
	if i < 0 {
		fmt.Println("Before panic")
		panic("Input must be a non-negative integer")
	}
	fmt.Println("Processing value:", i)
}
