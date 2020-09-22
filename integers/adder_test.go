package integers

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	sum := Add(3, 5)
	expected := 8

	if sum != expected {
		t.Errorf("got %d want %d", sum, expected)
	}
}

/*
 * Examples are compiled as part of test suite.
 * 1) Should start with Example, just like Test.
 * 2) Use the same format as below like fmt.
 * 3) To see example getting testing use
 *		go test -v
 * 4) Must use commented out //Output: value,
 *    else it won't compile
 *    It test output expected value with the fmt value.
 * 5) godoc will contain our defined functions with examples,
 *    which we defined below. (godoc -http=:6060)
 */
func ExampleAdd() {
	sum := Add(1, 2)
	fmt.Println(sum)
	// Output: 3
}
