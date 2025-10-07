package maps

import "errors"

type Dictionary map[string]string

// the internal reference to the map is copied to method
// technically, a copy, but point to the same location in memory
// changes here will change state
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", errors.New("word not found")
	}
	return definition, nil
}
