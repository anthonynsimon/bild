package convolution

import "testing"

func TestNewKernel(t *testing.T) {
	cases := []struct {
		size     int
		expected *Kernel
	}{
		{
			size:     0,
			expected: &Kernel{[]float64{}, 0, 0},
		},
		{
			size: 1,
			expected: &Kernel{[]float64{
				0,
				0,
			}, 1, 1},
		},
		{
			size: 2,
			expected: &Kernel{[]float64{
				0, 0,
				0, 0,
			}, 2, 2},
		},
		{
			size: 3,
			expected: &Kernel{[]float64{
				0, 0, 0,
				0, 0, 0,
				0, 0, 0,
			}, 3, 3},
		},
		{
			size: 10,
			expected: &Kernel{[]float64{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			}, 10, 10},
		},
	}

	for _, c := range cases {
		actual := NewKernel(c.size, c.size)
		if !kernelEqual(actual, c.expected) {
			t.Errorf("%s: expected: %#v, actual: %#v", "NewKernel", c.expected, actual)
		}
	}
}

func TestAbsum(t *testing.T) {
	cases := []struct {
		kernel   *Kernel
		expected float64
	}{
		{
			expected: 0,
			kernel:   NewKernel(0, 0),
		},
		{
			expected: 10,
			kernel: &Kernel{[]float64{
				5, 0, 1,
				0, 2, 0,
				0, 2, 0,
			}, 3, 3},
		},
		{
			expected: 11,
			kernel: &Kernel{[]float64{
				4, 0, 1,
				0, 1, 0,
				1, 3, 1,
			}, 3, 3},
		},
		{
			expected: 34,
			kernel: &Kernel{[]float64{
				20, 0, 2,
				0, -9, 0,
				-2, 0, -1,
			}, 3, 3},
		},
		{
			expected: 11,
			kernel: &Kernel{[]float64{
				0, 0, 0, -1,
				0, 9, 0, 0,
				0, 0, 0, -1,
				0, 0, 0, 0,
			}, 4, 4},
		},
	}

	for _, c := range cases {
		actual := c.kernel.Absum()
		if actual != c.expected {
			t.Errorf("%s: expected: %#v, actual: %#v", "KernelAbSum", c.expected, actual)
		}
	}
}

func TestKernelAt(t *testing.T) {
	cases := []struct {
		x, y     int
		kernel   *Kernel
		expected float64
	}{
		{
			x:        0,
			y:        0,
			expected: 5,
			kernel: &Kernel{[]float64{
				5, 0, 1,
				0, 2, 0,
				0, 2, 0,
			}, 3, 3},
		},
		{
			x:        2,
			y:        1,
			expected: -2,
			kernel: &Kernel{[]float64{
				4, -7, 1,
				-11, 1, -2,
				1, 3, 1,
			}, 3, 3},
		},
		{
			x:        2,
			y:        2,
			expected: -1,
			kernel: &Kernel{[]float64{
				20, 0, 2,
				0, -9, 0,
				-2, 0, -1,
			}, 3, 3},
		},
		{
			x:        3,
			y:        2,
			expected: -1,
			kernel: &Kernel{[]float64{
				0, 0, 0, -1,
				0, 9, 0, 0,
				0, 0, 0, -1,
				0, 0, 92, 0,
			}, 4, 4},
		},
	}

	for _, c := range cases {
		actual := c.kernel.At(c.x, c.y)
		if actual != c.expected {
			t.Errorf("%s: expected: %#v, actual: %#v", "KernelAt", c.expected, actual)
		}
	}
}

func TestKernelNormalized(t *testing.T) {
	cases := []struct {
		desc     string
		kernel   *Kernel
		expected *Kernel
	}{
		{
			desc: "all zero",
			kernel: &Kernel{[]float64{
				0, 0, 0,
				0, 0, 0,
				0, 0, 0,
			}, 3, 3},
			expected: &Kernel{[]float64{
				0, 0, 0,
				0, 0, 0,
				0, 0, 0,
			}, 3, 3},
		},
		{
			desc: "one element",
			kernel: &Kernel{[]float64{
				0, 0, 0,
				0, 1, 0,
				0, 0, 0,
			}, 3, 3},
			expected: &Kernel{[]float64{
				0, 0, 0,
				0, 1, 0,
				0, 0, 0,
			}, 3, 3},
		},
		{
			desc: "sum 3",
			kernel: &Kernel{[]float64{
				0, 0, 0,
				1, 1, 0,
				0, 0, 1,
			}, 3, 3},
			expected: &Kernel{[]float64{
				0, 0, 0,
				1.0 / 3, 1.0 / 3, 0,
				0, 0, 1.0 / 3,
			}, 3, 3},
		},
		{
			desc: "sum 4",
			kernel: &Kernel{[]float64{
				0, 0, 0,
				1, -2, 0,
				0, 0, 1,
			}, 3, 3},
			expected: &Kernel{[]float64{
				0, 0, 0,
				1.0 / 4, -2.0 / 4, 0,
				0, 0, 1.0 / 4,
			}, 3, 3},
		},
		{
			desc: "sum 5",
			kernel: &Kernel{[]float64{
				0, 0, 0,
				1, -2, 0,
				-1, 0, 1,
			}, 3, 3},
			expected: &Kernel{[]float64{
				0, 0, 0,
				1.0 / 5, -2.0 / 5, 0,
				-1.0 / 5, 0, 1.0 / 5,
			}, 3, 3},
		},
		{
			desc: "single negative element",
			kernel: &Kernel{[]float64{
				0, 0, 0,
				0, -1, 0,
				0, 0, 0,
			}, 3, 3},
			expected: &Kernel{[]float64{
				0, 0, 0,
				0, -1, 0,
				0, 0, 0,
			}, 3, 3},
		},
	}

	for _, c := range cases {
		actual := c.kernel.Normalized()
		if !kernelEqual(actual.(*Kernel), c.expected) {
			t.Errorf("%s: expected: %#v, actual: %#v", "KernelNormalized "+c.desc, c.expected, actual)
		}
	}
}

func TestKernelString(t *testing.T) {
	cases := []struct {
		kernel   *Kernel
		expected string
	}{
		{
			kernel: &Kernel{[]float64{
				0, 0, 0,
				0, -1, 0,
				0, 0, 0,
			}, 3, 3},
			expected: "\n0.0000  0.0000  0.0000  \n0.0000  -1.0000 0.0000  \n0.0000  0.0000  0.0000  ",
		},
		{
			kernel: &Kernel{[]float64{
				-2.75, 0, 0,
				0, -1, 0,
				0, 0, 92.32579,
			}, 3, 3},
			expected: "\n-2.7500 0.0000  0.0000  \n0.0000  -1.0000 0.0000  \n0.0000  0.0000  92.3258 ",
		},
	}

	for _, c := range cases {
		actual := c.kernel.String()
		if actual != c.expected {
			t.Errorf("%s: expected: %#v, actual: %#v", "KernelString", c.expected, actual)
		}
	}
}

func kernelEqual(a, b *Kernel) bool {
	if a.Matrix == nil && b.Matrix == nil {
		return true
	}

	if a.MaxX() != b.MaxX() || a.MaxY() != b.MaxY() {
		return false
	}

	for x := 0; x < a.MaxX(); x++ {
		for y := 0; y < a.MaxY(); y++ {
			if a.Matrix[y*a.MaxX()+x] != b.Matrix[y*b.MaxX()+x] {
				return false
			}
		}
	}

	return true
}
