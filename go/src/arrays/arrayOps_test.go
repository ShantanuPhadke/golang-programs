package arrays

import "testing"

func TestArraySum(t *testing.T) {
	// Declaring a generalized function for checking if the expected
	// and actual messages are equal or not
	assertCorrectValue := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("testing the array sum function", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		arraySum := ArraySum(numbers)
		expected := 15
		assertCorrectValue(t, arraySum, expected)
	})

	t.Run("testing the array sum function with an empty array", func(t *testing.T) {
		numbersEmpty := []int{}
		emptyArraySum := ArraySum(numbersEmpty)
		expected := 0
		assertCorrectValue(t, emptyArraySum, expected)
	})
}

func TestArraySumAll(t *testing.T) {
	// Declaring a generalized function for checking if the expected
	// and actual messages are equal or not
	assertCorrectArray := func(t *testing.T, got, want []int) {
		t.Helper()
		for index, sum := range want {
			if got[index] != sum {
				t.Errorf("sum at index %q was actually %q but was expected to be %q", index, got[index], sum)
			}
		}
	}

	t.Run("testing the ArraySumAll function on a basic case", func(t *testing.T) {
		sums := ArraySumAll([]int{1, 3}, []int{2, 3}, []int{9, 2})
		expectedArraySum := []int{4, 5, 11}
		assertCorrectArray(t, sums, expectedArraySum)
	})
}
