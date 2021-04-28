package main

// Define type Dictionary which acts as a wrapper around map
// map[key]value types are key string and value string
type Dictionary map[string]string

// Make errors constant. This requires us to create our own DictionaryErr type
// which implements the "error" interface.
// This makes errors more reusable and immutable.
const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

// By having Error() method, DictionaryErr implements the "error" interface
// DictionaryErr.Error() prints the string representation of error
func (e DictionaryErr) Error() string {
	return string(e)
}

// Returns word and error ie. string and error
// Looking up a value from a map can return 2 values.
// The value of key, and second value of boolean which indicates if the key was
// found successfully. If "false" no key exists
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// d.Search(word) returns definition and error. We are only interested in error
// so definition is omitted using black variable _
// Use switch statement for different errors
func (d Dictionary) Add(word string, definition string) error {
	_, err := d.Search(word)

	switch err {
	// No key was found for word. Add new key and value.
	case ErrNotFound:
		d[word] = definition
	// No error was returned. Key of word already exists
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word string, newDefinition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = newDefinition
	default:
		return err
	}
	return nil
}

// Go has a built-in function delete that workds on maps. It takes two arguments.
// The first is the map, and the second is the key to be removed
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
