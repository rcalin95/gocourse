package basics

import "fmt"

func main() {

		process()
		fmt.Println("Main function continues after recover.")

}

func process() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Start Process")
	panic("Something went wrong!")
	fmt.Println("End Process")
}
