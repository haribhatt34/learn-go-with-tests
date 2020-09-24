package arrays

import "testing"

// Test Coverage -
// go test -cover.
// if TDD is strictly followed, test coverage is near to 100%.

func TestSum(t *testing.T) {

	//array
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})

	//slice
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})
}
