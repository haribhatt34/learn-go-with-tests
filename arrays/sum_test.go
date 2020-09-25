package arrays

import (
	"reflect"
	"testing"
)

// Test Coverage -
// go test -cover.
// if TDD is strictly followed, test coverage is near to 100%.

func TestSum(t *testing.T) {

	//slice
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

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{2, 3})
	want := []int{6, 5}

	// we cannot compare two slices a, b as a == b.
	// we have to compare them one by one.
	// or we can also use reflect.DeepEqual but it is not type safe,
	// i.e. it will not throw error if we compare say string with int.

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	assertSumAllTails := func(got, want []int, t *testing.T) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("sum of non empty slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{2, 3})
		want := []int{5, 3}
		assertSumAllTails(got, want, t)
	})
	t.Run("sum of empty slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{})
		want := []int{5, 0}
		assertSumAllTails(got, want, t)
	})
}
