package main

import "testing"

func TestSearch(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		_, got := dictionary.Search("testing")
		want := errNotFound

		if got == nil {
			t.Fatal("expected error")
		}

		assertError(t, got, want)
	})

}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		got := dictionary.Add("test", "this is a test")
		assertError(t, got, nil)
		assertDef(t, dictionary, "test")
	})

	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is a test"}
		got := dictionary.Add("test", "this is a test")
		assertError(t, got, errAlreadyExists)
		assertDef(t, dictionary, "test")
	})
}

func TestUpdate(t *testing.T) {
	t.Run("word exists", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a word"}
		err := dictionary.Update("test", "this is a test")

		assertError(t, err, nil)
		assertDef(t, dictionary, "test")
	})

	t.Run("word does not exists", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Update("test", "this is a test")

		assertError(t, err, errAlreadyExists)
	})
}

func TestDelete(t *testing.T) {
	t.Run("word exists", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a word"}
		err := dictionary.Delete("test")
		_, searchErr := dictionary.Search("test")

		assertError(t, err, nil)
		assertError(t, searchErr, errNotFound)
	})

	t.Run("word does not exists", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Delete("test")

		assertError(t, err, errNotFound)
	})
}

func assertDef(t testing.TB, dictionary Dictionary, key string) {
	t.Helper()

	got, _ := dictionary.Search(key)
	want := "this is a test"

	assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("wanted %q, but got %q", want, got)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("wanted %q, but got %q", want, got)
	}
}
