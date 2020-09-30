package maps

// Search function returns the value for a given key
func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}
