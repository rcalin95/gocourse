package intermediate

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	message := "Hello, \nGo!"
	message1 := "Hello, \tGo!"
	message2 := "Hello, Go!" // Go!lo 
	rawMessage := `Hello\nGo`

	fmt.Println(message)
	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(rawMessage)

	fmt.Println("Length of rawMessage variable is", len(rawMessage))
	fmt.Println("The first character in message var is ", message[0]) // ASCII

	greeting := "Hello "
	name := "Alice"
	fmt.Println(greeting + name)

	str1 := "Apple"
	str := "apple"
	str2 := "Banana"
	str3 := "App"
	fmt.Println(str1 < str2) // comparison is done on ASCII values
	fmt.Println(str3 < str1) // and is done character by character
	fmt.Println(str > str1)

	for i, char := range message {
		fmt.Printf("Character at index %d is %c\n",i,char)
		fmt.Printf("%v\n", char)
	}

	fmt.Println("Rune count:", utf8.RuneCountInString(greeting))

	greetingWithName := greeting + name
	fmt.Println(greetingWithName)

	var ch rune = 'a'

	fmt.Printf("%c\n", ch)

	cstr := string(ch)
	fmt.Println(cstr)

	fmt.Printf("Type of cstr is %T\n", cstr)

	myString := "Hello, Golang!"
	for _, v := range myString{
		fmt.Printf("%c\n", v)
	}
}
