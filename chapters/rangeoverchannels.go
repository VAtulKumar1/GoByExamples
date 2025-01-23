package chapters

import "fmt"

func RangeOverChannels() {
	q := make(chan string, 2)
	q <- "one"
	q <- "two"
	close(q)

	for elem := range q {
		fmt.Println(elem)
	}
}
