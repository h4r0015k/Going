package array

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Array of size 5", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("want %d but got %d, %v", want, got, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{1, 3, 4})
	want := []int{3, 8}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %d but got %d", want, got)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSum := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %d but got %d", want, got)
		}
	}

	t.Run("Sum tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{2, 3, 3})
		want := []int{2, 6}

		checkSum(t, got, want)
	})

	t.Run("Safely sum tails with empty tails", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{2, 3, 3})
		want := []int{0, 6}

		checkSum(t, got, want)
	})
}
