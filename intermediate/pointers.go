package intermediate

import "fmt"

func main() {

	// var ptr *int
	var ptr *int
	var a int = 10
	ptr = &a // Referencing a pointer

	fmt.Println(a)
	// fmt.Println(ptr) // 0xc0000aa008
	// fmt.Println(*ptr) // derefencing a pointer

	if ptr == nil {
		fmt.Println("Pointer is nil")
	}

	modifyValue(ptr)
	fmt.Println(a)
}

func modifyValue (ptr *int){	
	*ptr++
}
