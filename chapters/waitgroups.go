package chapters

import (
	"fmt"
	"sync"
	"time"
)

func Workers(id int) {
	fmt.Printf("worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("worker %d finished\n", id)
}

func WaitGroups() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			Workers(i)
		}()
	}

	wg.Wait()
}
