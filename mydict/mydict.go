package mydict

import "errors"

// Dictionary type
// Go는 struct 뿐만 아닌 type에도 메서드를 정의할 수 있음.
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")
var errWordExists = errors.New("That Word already existed")
var errUpdate = errors.New("Updating Error gg!")

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dict
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

// Update a word to the dict
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	// 단어가 존재한다면,
	if err == nil {
		d[word] = def
	} else if err == errNotFound {
		return errUpdate
	}
	return nil
}

// Delete a word to the dict
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
