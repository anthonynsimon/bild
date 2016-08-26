package f64

import "testing"

func TestClamp(t *testing.T) {
	cases := []struct {
		value, min, max, expected float64
	}{
		{
			value:    0.0,
			min:      0.0,
			max:      1.0,
			expected: 0.0,
		},
		{
			value:    0.0,
			min:      1.0,
			max:      1.5,
			expected: 1.0,
		},
		{
			value:    200.0,
			min:      0.0,
			max:      1.0,
			expected: 1.0,
		},
		{
			value:    -4561200.0,
			min:      0.0,
			max:      1.0,
			expected: 0.0,
		},
	}

	for _, c := range cases {
		actual := Clamp(c.value, c.min, c.max)
		if actual != c.expected {
			t.Errorf("f64.Clamp: expected: %v actual: %v", c.expected, actual)
		}
	}
}
