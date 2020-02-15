package main

// CHANGE LOG:
// (1) Started off with just a simple Search function that takes in
// 	   a basic string to string map and a searchKey string as inputs
// (2) Made a Dictionary type and put the Search function in as one
//     of its methods.
// (3) Added support for returning error when an unknown searchKey is
//     supplied to the Search method.
// (4) Made a new function that just takes in Dictionary Error and returns
//     the associated string for the inputted error.
// (5) Added extra Update and Delete methods.

/*Dictionary ... type of map string to string */
type Dictionary map[string]string

/*DictionaryErr ... string representing type of error */
type DictionaryErr string

const (
	/*ErrorSearchKeyNotFound ... Error for when an user tries to search for an invalid key*/
	ErrorSearchKeyNotFound = DictionaryErr("supplied search key doesn't exist in the dictionary")
	/*ErrorAddDuplicateKey ... Error for when user tries to insert an already existing key*/
	ErrorAddDuplicateKey = DictionaryErr("cannot add a duplicate key to a dictionary without deleting the original entry")
	/*ErrorUpdateKeyNotFound ... Error for when user tries to update a key that isn't in the dictionary*/
	ErrorUpdateKeyNotFound = DictionaryErr("cannot update a key that isn't in the dictionary")
	/*ErrorDeleteKeyNotFound ... Error for when user tries to delete a (key, value) pair where the key isn't in the dictionary*/
	ErrorDeleteKeyNotFound = DictionaryErr("cannot update a key that isn't in the dictionary")
)

/*Error ... function for returning associated String for an inputted DictionaryErr*/
func (e DictionaryErr) Error() string {
	return string(e)
}

/*Search ... function for searching a dictionary */
func (dictionary Dictionary) Search(searchKey string) (string, error) {
	value, searchKeyValid := dictionary[searchKey]
	if !searchKeyValid {
		return "", ErrorSearchKeyNotFound
	}
	return value, nil
}

/*Add ... function for adding a new entry to a dictionary*/
func (dictionary Dictionary) Add(dictionaryKey string, dictionaryValue string) (bool, error) {
	_, keyValid := dictionary[dictionaryKey]
	if !keyValid {
		// Insert the new key, value pair
		return true, nil
	}
	// If the key already exists in the dictionary just return false and the error
	return false, ErrorAddDuplicateKey
}

/*Update ... function for updating a (key, value) pair in a dictionary*/
func (dictionary Dictionary) Update(updateKey string, newValue string) (bool, error) {
	_, keyValid := dictionary[updateKey]
	if !keyValid {
		// Key is not in the dictionary
		return false, ErrorUpdateKeyNotFound
	}
	// Key is in the dictionary
	dictionary[updateKey] = newValue
	return true, nil
}

/*Delete ... function for deleting an entry in a dictionary*/
func (dictionary Dictionary) Delete(deleteKey string) (bool, error) {
	_, keyValid := dictionary[deleteKey]
	if !keyValid {
		// Key is not in the dictionary
		return false, ErrorDeleteKeyNotFound
	}
	// Key is in the dictionary
	delete(dictionary, deleteKey)
	return true, nil
}

func main() {

}
