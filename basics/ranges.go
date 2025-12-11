package basics

import "fmt"

func main() {

	message := "Hello World!"

	for _, v := range message {
		// fmt.Println(i, v)
		fmt.Printf("Character: %c\n", v)
	}


}
