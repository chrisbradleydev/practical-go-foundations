package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	/* solution 1: use mutext
	var mu sync.Mutex // mutex should be ontop of guarded variable(s)
	count := 0
	*/
	/* solution 2: use atomic
	atomic functions require great care to be used correctly
	except for special, low-level applications, synchronization
	is better done with channels or the facilities of the sync package
	*/
	// count := int64(0)
	var count int64

	var wg sync.WaitGroup
	const n = 10
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10_000; j++ {
				/* solution 1: use mutext
				mu.Lock()
				count++
				mu.Unlock()
				*/
				// solution 2: use atomic
				atomic.AddInt64(&count, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println(count)
}