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
