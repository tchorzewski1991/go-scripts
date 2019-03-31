package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// note is a simple struct that embeds io.Reader interface.
// I want to state explicitly that the embedding struct needs
// to satisfy the embedded interface and at the same time hide
// it’s data. Another reason is that I want to get access to
// Read() method which already has been implemented in several
// places. One of the examples could be *strings.Reader type.
// I want my own struct, but there is no need to reinvent the
// wheel - at least, not in such basic case.
type note struct {
	text string
	io.Reader
}

// newNote returns new *note reading from s. It uses strings.NewReader()
// to inject object, that satisfies io.Reader interface.
func newNote(s string) *note {
	return &note{s, strings.NewReader(s)}
}

// reverseNote is a simple struct that will implement custom Read() method
// to satisfy io.Reader interface manually.
type reverseNote struct {
	text string
	i    int64
}

// newReverseNote returns new *ReverseNote reading from s.
func newReverseNote(s string) *reverseNote {
	return &reverseNote{s, 0}
}

// Read satisfies io.Reader interface for ReversNote type.
// Internally this method takes each character from underlying
// string source and writes it to provided slice of bytes in
// reverse order.
func (rn *reverseNote) Read(p []byte) (n int, err error) {
	if rn.i >= int64(len(rn.text)) {
		return 0, io.EOF
	}

	k := len(rn.text) - 1
	var temp = make([]byte, k+1)

	for i := 0; i <= k; i++ {
		temp[i] = rn.text[k-i]
	}

	n = copy(p, temp)
	rn.i += int64(n)

	return
}

func main() {
	n := newNote("Text")
	rn := newReverseNote("Text")

	displayNote(n)
	fmt.Println("\n")
	displayNote(rn)
}

// displayNote takes anything that implements io.Reader interface and
// delegates it internally to io.Copy. They way displayNote() is implemented
// does not matter much. The main point here was to get hands on even more
// practise how interfaces works as well as how to implement your own
// reading logic.
// - First case uses standard library implementation of Read() method.
//   Internally it uses *strings.Reader Read() method. It should print "Text".
// - Second case is the custom one. It reads given source string in
//   reverse order. It should print "txeT".
func displayNote(r io.Reader) {
	io.Copy(os.Stdout, r)
}
