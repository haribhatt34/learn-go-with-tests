package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	// Struct Buffer is variable size buffer of bytes with read/write access
	// and ready to use.
	buffer := bytes.Buffer{}
	Greet(&buffer, "Hari")

	// String return the content of the buffer as string.
	got := buffer.String()
	want := "Hello, Hari"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

/* Notes:
 *
 * Dependency Injection(to pass in) here means -
 *
 * We are basically trying to pass the dependency of the printing,
 * from our test functions so that we can test the actually printing.
 *
 * Below is the flow of Printf
 * func Greet(name string) {
 * 		fmt.Printf("Hello, %s", name)
 * }
 *
 * // Printf return the no of bytes written and any error encountered.
 * func Printf(format string, a ...interface{}) (n int, err error) {
 * 		return Fprintf(os.stdout, format, a...)
 * }
 *
 * func Fprintf(w io.Writer, format string, a ...interface) (n int, err error) {
 * 		p := newPrinter()
 * 		p.doPrintf(format, a)
 * 		n, err = w.Write(p.buf)
 * 		p.free()
 *		return
 * }
 *
 * An io.Writer
 * type Writer interface {
 * 		Write(p []byte) (n int, err error)
 * }
 */
