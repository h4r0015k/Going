package main

import "errors"

type Dictionary map[string]string

var errNotFound = errors.New("not found")
var errAlreadyExists = errors.New("key already exists")

func (d Dictionary) Search(word string) (string, error) {
	value, found := d[word]

	if !found {
		return "", errNotFound
	}

	return value, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	if err == nil {
		return errAlreadyExists
	}

	d[key] = value
	return nil
}

func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}
