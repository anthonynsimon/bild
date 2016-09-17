package util

import (
	"testing"
)

func TestStack(t *testing.T) {
	var st Stack

	floatElement := 12.0
	stringElement := "el"
	intElement := 3

	st.Push(floatElement)
	st.Push(stringElement)
	st.Push(intElement)

	if 3 != st.Len() {
		t.Fatalf("Expected stack length of %d, got %d", 3, st.Len())
	}

	pop1 := st.Pop()
	pop2 := st.Pop()
	pop3 := st.Pop()
	noElementsLeft := 0 == st.Len()

	if intElement != pop1 {
		t.Fatalf("Expected element to be %d, got %d", intElement, pop1)
	}
	if stringElement != pop2 {
		t.Fatalf("Expected element to be %s, got %s", stringElement, pop2)
	}
	if floatElement != pop3 {
		t.Fatalf("Expected element to be %6.2f, got %6.2f", floatElement, pop3)
	}
	if !noElementsLeft {
		t.Fatalf("Expected noElementsLeft to be %t, got %t", true, noElementsLeft)
	}
}
