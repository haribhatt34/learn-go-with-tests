package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Greet says "Hello, name"
// We aret telling Fprintf to write data in writer object.
// instead of usual os.stdout from fprintf
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// MyGreeterHandler passes data to Greet function to print on Browser
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "World")
}

func main() {
	Greet(os.Stdout, "Bhatt")
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}

/* Notes
 * 1) To change the Greet function to more general, we can change
 * 1st argument from writer *bytes.Buffer to
 * 					 writer *io.Writer
 * 2) Through this generic Implementation of Greet, we can write
 * date into -
 * i)   byte.Buffer
 * ii)  os.stdout
 * iii) internet
 */
