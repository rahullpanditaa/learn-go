package maps

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("word not found")

// the internal reference to the map is copied to method
// technically, a copy, but point to the same location in memory
// changes here will change state
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}
