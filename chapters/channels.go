package chapters

import "fmt"

func Channel() {
	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)

	msg2 := make(chan string, 2)

	msg2 <- "bufferd"
	msg2 <- "channel"

	fmt.Println(<-msg2)
	fmt.Println(<-msg2)
}
