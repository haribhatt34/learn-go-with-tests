package iteration

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("a")
	want := "aaaaa"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

/* Benchmark inbuilt function inside package testing
 * Benchmark, Example, Testing has all have similar syntax.
 * to run benchmark on console type :
 * go test -bench=.
 * 
 * By default benchmark are run sequentially.
 */
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
