package maps

import (
	"errors"
)

// Dictionary acts as a wrapper around map
type Dictionary map[string]string

// ErrNotFound contains the message when
var ErrNotFound = errors.New("This word doesn't exist in dictionary")

// Search method returns the value for a given key
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// Add method add given word and its meaning to the dictionary
func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}
