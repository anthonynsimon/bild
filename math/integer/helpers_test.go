package integer

import "testing"

func TestMin(t *testing.T) {
	cases := []struct {
		a, b, expected int
	}{
		{
			a:        0,
			b:        0,
			expected: 0,
		},
		{
			a:        1,
			b:        1,
			expected: 1,
		},
		{
			a:        -1,
			b:        1,
			expected: -1,
		},
		{
			a:        1,
			b:        -1,
			expected: -1,
		},
		{
			a:        10,
			b:        2,
			expected: 2,
		},
	}

	for _, c := range cases {
		actual := Min(c.a, c.b)
		if actual != c.expected {
			t.Errorf("Min: expected: %v actual: %v", c.expected, actual)
		}
	}
}

func TestMax(t *testing.T) {
	cases := []struct {
		a, b, expected int
	}{
		{
			a:        0,
			b:        0,
			expected: 0,
		},
		{
			a:        1,
			b:        1,
			expected: 1,
		},
		{
			a:        -1,
			b:        1,
			expected: 1,
		},
		{
			a:        1,
			b:        -1,
			expected: 1,
		},
		{
			a:        10,
			b:        2,
			expected: 10,
		},
	}

	for _, c := range cases {
		actual := Max(c.a, c.b)
		if actual != c.expected {
			t.Errorf("Max: expected: %v actual: %v", c.expected, actual)
		}
	}
}
