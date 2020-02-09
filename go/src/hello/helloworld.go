package main

import "fmt"

// LOG OF CHANGES
// (1) Separating out what we want printed form the actual printing
// (2) Added an extra parameter to the Hello function as a greeting
// (3) Refactored the string "Hello world " out of literal and into a constant
// (4) Made the Hello function to support two cases
// 		(a) Return "Hello, world" when name is empty string
// 		(b) Return "Hello world, <name>" when name is NOT empty string
// (5) Changing Hello to take in a second parameter that specifies the language of the return message
//		(a) Default to English when the language parameter passed in is not valid
// (6) Make the Hello function more concise by putting extra functionality into a second function for forming
// 	the prefix we will end up returning

const englishHelloPrefix = "Hello"
const englishWorldPrefix = " world, "

const spanishHelloPrefix = "Hola"
const spanishWorldPrefix = " mundo, "

const frenchHelloPrefix = "Bonjour"
const frenchWorldPrefix = " monde, "

// Defining our language constants below

const spanish = "Spanish"
const french = "French"

/*Hello ... Function Specification:
INPUT(S):
	name = A string carrying the name of the person we want to greet
		   otherwise can be an empty string
	language = The language in which we want to retunr our greeting in
OUTPUT(S):
	Return "Hello world, <name>" if name is a valid string
	Otherwise return "Hello, world no name!" if name is empty string
*/
func Hello(name string, language string) string {
	defaultName := name
	if name == "" {
		switch language {
		case spanish:
			defaultName = ", mundo"
		case french:
			defaultName = ", monde"
		default:
			defaultName = ", world"
		}
	}
	return greetingPrefix(name, language) + defaultName
}

/*greetingPrefix ... Function Specification:
INPUT(S):
	name = A string carrying the name of the person we want to greet
		   otherwise can be an empty string
	language = The language in which we want to retunr our greeting in
OUTPUT(S):
	Returns the appropriate prefix for our Hello method given the
	name and language inputs

*/
func greetingPrefix(name string, language string) (prefix string) {
	if name == "" {
		switch language {
		case spanish:
			prefix = spanishHelloPrefix
		case french:
			prefix = frenchHelloPrefix
		default:
			prefix = englishHelloPrefix
		}
	} else {
		switch language {
		case spanish:
			prefix = spanishHelloPrefix + spanishWorldPrefix
		case french:
			prefix = frenchHelloPrefix + frenchWorldPrefix
		default:
			prefix = englishHelloPrefix + englishWorldPrefix
		}
	}
	return
}

func main() {
	fmt.Println(Hello("Shantanu", ""))
}
