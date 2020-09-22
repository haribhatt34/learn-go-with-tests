package integers

import "testing"

func TestSum(t *testing.T) {
	sum := Add(3, 5)
	expected := 8

	if sum != expected {
		t.Errorf("got %d want %d", sum, expected)
	}
}
