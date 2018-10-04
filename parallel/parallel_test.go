package parallel

import "testing"

func TestParallelize(t *testing.T) {
	for n := 0; n < 1024; n++ {
		data := make([]bool, n)

		Line(len(data), func(start, end int) {
			for i := start; i < end; i++ {
				data[i] = !data[i]
			}
		})

		for _, d := range data {
			if !d {
				t.Errorf("Test parallelize failed. Failure at n = %v", n)
			}
		}
	}
}
