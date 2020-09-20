package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func Benchmark(b *testing.B) {
	for i := 1; i <= b.N; i++ {
		Hello("Chris")
	}
}
