/*Package parallel provides helper functions for the dispatching of parallel jobs.*/
package parallel

import (
	"runtime"
	"sync"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// Line dispatches a parameter fn into multiple goroutines by splitting the parameter length
// by the number of available CPUs and assigning the length parts into each fn.
func Line(length int, fn func(start, end int)) {
	procs := runtime.GOMAXPROCS(0)
	counter := length
	partSize := length / procs
	if procs <= 1 || partSize <= procs {
		fn(0, length)
	} else {
		var wg sync.WaitGroup
		for counter > 0 {
			start := counter - partSize
			end := counter
			if start < 0 {
				start = 0
			}
			counter -= partSize
			wg.Add(1)
			go func() {
				defer wg.Done()
				fn(start, end)
			}()
		}

		wg.Wait()
	}
}
