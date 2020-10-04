package maps

// Dictionary acts as a wrapper around map
type Dictionary map[string]string

// DictionaryErr acts as a wrapper around errors.New()
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	// ErrNotFound is returned when search can't find given word in dictionary.
	ErrNotFound = DictionaryErr("This word doesn't exist in dictionary")

	// ErrWordExists returned when we try to add a word which already exists in the dictionary.
	ErrWordExists = DictionaryErr("This word already exist in the dictionary")
)

// Search method returns the value for a given key
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// Add method add given word and its meaning to the dictionary
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
		// this case is added if there is any other error except for defined one(i.e. ErrNotFound)
	default:
		return err
	}

	return nil
}

// Notes:
/*
 * 1) We cannot use errors.New("") with const.
 * 2) Earlier we were using below
 * 	var ErrWordExists = errors.New("This word already exist in the dictionary")
 *  But we cannot use it with const instead of var.
 *  So, with const we are required to create our own error type DictionaryErr.
 *  The dictionary type implements error interface.
 *
 */
