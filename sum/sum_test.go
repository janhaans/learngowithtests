package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("sum a slice of any numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("Got %v, want %v, given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("Sum all slices", func(t *testing.T) {
		s1 := []int{1, 2, 3, 4, 5}
		s2 := []int{1, 2, 3}
		got := SumAll(s1, s2)
		want := []int{15, 6}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v, want %v", got, want)
		}

	})
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v, want %v", got, want)
		}
	}

	t.Run("Sum all slice tails", func(t *testing.T) {
		s1 := []int{1, 2, 3, 4, 5}
		s2 := []int{1, 2, 3}
		got := SumAllTails(s1, s2)
		want := []int{14, 5}

		checkSums(t, got, want)
	})

	t.Run("Sum all slice tails with empty slice", func(t *testing.T) {
		s1 := []int{}
		s2 := []int{1, 2, 3}
		got := SumAllTails(s1, s2)
		want := []int{0, 5}

		checkSums(t, got, want)
	})
}
