package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("SYS: Cannot found the word.")
var errWordExists = errors.New("SYS: Word already exists.")

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}
