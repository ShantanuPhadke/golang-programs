package main

import "testing"

func TestHello(t *testing.T) {

	// Declaring a generalized function for checking if the expected
	// and actual messages are equal or not
	assertCorrectMessage := func(t *testing.T, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying Hello world to people", func(t *testing.T) {
		got := Hello("Shantanu", "")
		want := "Hello world, Shantanu"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to world", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("test the Hello function in Spanish", func(t *testing.T) {
		got := Hello("Shantanu", "Spanish")
		want := "Hola mundo, Shantanu"
		assertCorrectMessage(t, got, want)
	})

}
