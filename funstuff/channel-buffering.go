package funstuff

import "fmt"

func main() {
	
	messages := make(chan string, 2) // buffered channel with capacity of 2

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages) // This will block since the channel is empty
}