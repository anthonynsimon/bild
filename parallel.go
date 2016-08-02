package bild

import (
	"runtime"
	"sync"
)

// ParallelizationEnabled sets if package should use goroutines when possible
var ParallelizationEnabled = true

func parallelize(size int, fn func(start, end int)) {
	if !ParallelizationEnabled {
		fn(0, size)
	} else {
		var wg sync.WaitGroup
		procs := runtime.GOMAXPROCS(0)

		counter := size
		partSize := size / procs
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
