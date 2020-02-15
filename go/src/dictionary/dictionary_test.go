package main

import "testing"

func TestSearch(t *testing.T) {
	t.Run("testing search for a valid string", func(t *testing.T) {
		stringDictionary := Dictionary{"test key": "test value"}
		actualValue, _ := stringDictionary.Search("test key")
		expectedValue := "test value"
		assertStrings(t, actualValue, expectedValue)
	})

	t.Run("testing Search with an invalid string", func(t *testing.T) {
		stringDictionary := Dictionary{"test key": "test value"}
		_, err := stringDictionary.Search("invalid key")
		assertError(t, err, ErrorSearchKeyNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("testing adding in a valid key-value pair", func(t *testing.T) {
		stringDictionary := Dictionary{}
		addSuccessfulResult, _ := stringDictionary.Add("sample key", "sample value")
		expectedResult := true
		assertBoolean(t, addSuccessfulResult, expectedResult)
	})

	t.Run("testing adding in a duplicate key-value pair", func(t *testing.T) {
		stringDictionary := Dictionary{"sample key": "sample value"}
		unsuccessfulResult, err := stringDictionary.Add("sample key", "new sample value")
		expectedResult := false
		assertBoolean(t, unsuccessfulResult, expectedResult)
		assertError(t, err, ErrorAddDuplicateKey)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("testing a sample valid update", func(t *testing.T) {
		stringDictionary := Dictionary{"sample key": "sample value"}
		updateSuccessfulResult, _ := stringDictionary.Update("sample key", "new sample value")
		expectedResult := true
		assertBoolean(t, updateSuccessfulResult, expectedResult)
		assertStrings(t, stringDictionary["sample key"], "new sample value")
	})

	t.Run("testing an invalid update", func(t *testing.T) {
		stringDictionary := Dictionary{"sample key": "sample value"}
		updateUnsuccessfulResult, err := stringDictionary.Update("invalid key", "new sample value")
		expectedResult := false
		assertBoolean(t, updateUnsuccessfulResult, expectedResult)
		assertError(t, err, ErrorUpdateKeyNotFound)
	})
}

func TestDelete(t *testing.T) {
	t.Run("testing a sample valid delete", func(t *testing.T) {
		stringDictionary := Dictionary{"sample key": "sample value"}
		deleteSuccessfulResult, _ := stringDictionary.Delete("sample key")
		expectedResult := true
		assertBoolean(t, deleteSuccessfulResult, expectedResult)
		assertKeyNotInDictionary(t, stringDictionary, "sample key")
	})

	t.Run("testing a sample invalid delete", func(t *testing.T) {
		stringDictionary := Dictionary{"sample key": "sample value"}
		deleteUnsuccessfulResult, err := stringDictionary.Delete("sample key 2")
		expectedResult := false
		assertBoolean(t, deleteUnsuccessfulResult, expectedResult)
		assertError(t, err, ErrorDeleteKeyNotFound)
	})
}

func assertStrings(t *testing.T, actual, expected string) {
	t.Helper()
	if actual != expected {
		t.Errorf("Expected %q but got %q", expected, actual)
	}
}

func assertError(t *testing.T, err, expectedError error) {
	if err == nil {
		t.Errorf("No error returned when error expected")
	}
	if err != expectedError {
		t.Errorf("Wrong error returned for invalid key in Dictionary!")
	}
}

func assertBoolean(t *testing.T, actualBool, expectedBool bool) {
	t.Helper()
	if actualBool != expectedBool {
		t.Errorf("Expected %t but got %t", expectedBool, actualBool)
	}
}

func assertKeyNotInDictionary(t *testing.T, dictionary Dictionary, key string) {
	t.Helper()
	_, err := dictionary.Search(key)
	if err != ErrorSearchKeyNotFound {
		t.Errorf("The inputted key wasn't supposed to be in the inputted map but was")
	}
}
