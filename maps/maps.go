package main

type Dictionary map[string]string

var errNotFound = DictionaryErr("not found")
var errAlreadyExists = DictionaryErr("key already exists")

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	value, found := d[word]

	if !found {
		return "", errNotFound
	}

	return value, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case errNotFound:
		d[key] = value
	case nil:
		return errAlreadyExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case errNotFound:
		return errAlreadyExists
	case nil:
		d[key] = value
		return nil
	default:
		return err
	}
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)

	switch err {
	case errNotFound:
		return errNotFound
	case nil:
		delete(d, key)
		return nil
	default:
		return err
	}
}

func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}
