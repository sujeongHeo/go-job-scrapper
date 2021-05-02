package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var (
	errorNotFound = errors.New("Not Found")
	errCantUpdate = errors.New("Cant update non-exsisting word")
	errWordExists = errors.New("That word already exists")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]

	if exists {
		return value, nil
	}

	return "", errorNotFound

}

//Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errorNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}

//Update a word to the dictionary
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errorNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case errorNotFound:
		return errorNotFound
	case nil:
		delete(d, word)
	}
	return nil
}
