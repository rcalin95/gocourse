package basics

import "fmt"

var middleName = "Cane"

func main() {
	// var age int
	// var name string = "John"
	// var name = "Jane"

	// count := 10
	// lastName := "Smith"
	// middleName := "Cane"
	// Default values
	// Numeric types: 0
	// Boolean: false
	// String: ""
	// Pointers, slices, maps, functions, and structs: nil

	// -----SCOPE
	// middleName := "Changed"
	fmt.Println(middleName)
}

func printname(){
	firstName := "Michael"
	fmt.Println(firstName)
}
