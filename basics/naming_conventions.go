package basics

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// PascalCase
	// Eg. CalculateArea, UserInfo, NewHTTPRequest
	// Structs, Interfaces, enums

	// snake_case
	// Eg. user_id, first_name, http_request
	// variables

	// UPPERCASE
	// Use case is constants

	const MAXRETRIES = 5
	// mixedCase (camelCase)
	// Eg. javaScript, htmlDocument, isValid

	var employeeID = 1001
	fmt.Println("Employee ID: ", employeeID)
}
