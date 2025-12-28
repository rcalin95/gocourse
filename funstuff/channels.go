package funstuff

import "fmt"

func main() {

	messages := make(chan string)

	go func() { messages <- "ping" }() // anonymous goroutine sending "ping" to channel

	msg := <-messages // receiving from channel

	fmt.Println(msg)

}
