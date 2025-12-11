package basics

import (
	"fmt"
	"os"
)

func main() {

	defer fmt.Println("Defered statement")

	fmt.Println("Starting the main function")

	// Exit with status code 1
	os.Exit(1)

	//This line will not be executed
	fmt.Println("This line will not be printed")
}
