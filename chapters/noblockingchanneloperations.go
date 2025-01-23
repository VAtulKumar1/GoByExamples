package chapters

import "fmt"

func NonBlockingChannelOperations() {
	message := make(chan string)
	signals := make(chan string)

	select {
	case msg := <-message:
		fmt.Println("recieved message: ", msg)
	default:
		fmt.Println("no message recieved")
	}

	msg := "hi"

	select {
	case message <- msg:
		fmt.Println("Send Message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-message:
		fmt.Println("recieved message", msg)
	case sig := <-signals:
		fmt.Println("reccived signal ", sig)
	default:
		fmt.Println("no activilty")
	}
}
