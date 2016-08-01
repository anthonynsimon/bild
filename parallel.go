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
		partSize := size / procs
		wg.Add(procs)

		for i := 0; i < procs; i++ {
			i := i // needed for the closure to work
			go func() {
				defer wg.Done()
				fn(partSize*i, partSize*(i+1))
			}()
		}

		wg.Wait()
	}
}
