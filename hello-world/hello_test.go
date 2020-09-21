package main

import "testing"

/* In go we can declare functions inside other functions and assign them to variable.
* We pass t to assertCorrectMessage to tell test code to fail when we need to.
* t.Helper() tell the test suite that this method is a helper. So,
* when test fail, it will point to the actual sub-test line where error occured.
* to see this try to fail the test case.
 */

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	// below are subtests.
	// syntax (name, func)
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Chris", "Spanish")
		want := "Hola, Chris"
		assertCorrectMessage(t, got, want)
	})
}
