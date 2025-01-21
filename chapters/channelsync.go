package chapters

import (
	"fmt"
	"time"
)

func Worker(done chan bool) {
	fmt.Printf("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}
