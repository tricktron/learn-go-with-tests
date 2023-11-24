package maps

import "errors"

type (
	Dictionary      map[string]string
	DictionaryError string
)

const (
	ErrWordNotFound     = DictionaryError("could not find the word you were looking for")
	ErrWordExists       = DictionaryError("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryError("cannot update word because it does not exist")
)

func (e DictionaryError) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrWordNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	if err == nil {
		return ErrWordExists
	}

	d[word] = definition

	return nil
}

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)

	if errors.Is(err, ErrWordNotFound) {
		return ErrWordDoesNotExist
	}

	d[word] = newDefinition

	return nil
}
