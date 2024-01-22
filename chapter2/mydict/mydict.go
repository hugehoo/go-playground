package mydict

import (
	"errors"
	"fmt"
)

// Dictionary type
type Dictionary map[string]string

// when value is not exists
var (
	errNotFound   = errors.New("not found")
	errCantUpdate = errors.New("Can't Update Non-existing word")
	errWordExists = errors.New("That word already exists")
)

type NotFoundError struct {
	Word string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s not found", e.Word)
}

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	//return "", errNotFound
	return "", &NotFoundError{Word: word} //  creating a new NotFoundError object and obtaining its address in memory.
}

// Add a word to Dictionary
func (d Dictionary) Add(word string, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}

// Update a word, (why this arg doesn't require for asterik?)
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
