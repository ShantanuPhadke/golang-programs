package integers

import (
	"testing"
)

func TestOperations(t *testing.T) {

	// Declaring a generalized function for checking if the expected
	// and actual messages are equal or not
	assertCorrectValue := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("testing the add function", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4
		assertCorrectValue(t, sum, expected)
	})

	t.Run("testing the subtraction function", func(t *testing.T) {
		difference := Subtract(4, 2)
		expected := 2
		assertCorrectValue(t, difference, expected)
	})

	t.Run("testing the multiplication function", func(t *testing.T) {
		product := Multiply(4, 2)
		expected := 8
		assertCorrectValue(t, product, expected)
	})

	t.Run("testing the function", func(t *testing.T) {
		divisionResult := Divide(10, 2)
		expected := 5
		assertCorrectValue(t, divisionResult, expected)
	})
}
