package main

import "errors"

type Dictionary map[string]string

var errNotFound = errors.New("not found")

func (d Dictionary) Search(word string) (string, error) {
	value, found := d[word]

	if !found {
		return "", errNotFound
	}

	return value, nil
}

func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}
