package dictionary

import "errors"

type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not Found")
	errCantUpdate = errors.New("Can't update non-existing word")
	errWordExists = errors.New("Word already exists")
)

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]

	if exists {
		return value, nil
	}

	return "", errNotFound
}

func (d Dictionary) Add(word string, def string) error {
	_, err := d.Search(word)

	if err == nil {
		return errWordExists
	}

	d[word] = def

	return nil
}

func (d Dictionary) Update(word string, def string) error {
	_, err := d.Search(word)

	if err == nil {
		d[word] = def

		return nil
	}

	return errCantUpdate
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
