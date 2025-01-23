package chapters

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func AtomicCounters() {

	var ops atomic.Uint64

	var wg sync.WaitGroup

	for i := 1; i <= 50; i++ {
		wg.Add(1)

		go func() {
			for c := 1; c <= 1000; c++ {
				ops.Add(1)
			}

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("ops: ", ops.Load())
}
