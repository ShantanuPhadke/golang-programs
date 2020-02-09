package iteration

/*Repeat ... Function Specification:
INPUTS:
	character = Single character string which we want repeated the specified number of times
	repeatFactor = Integer specifying the number of times we want the character repeated
OUTPUTS:
	A string which contains the inputted character integer repeated the number of times specified
	via the input
*/
func Repeat(character string, repeatFactor int) string {
	var repeatedString string
	for i := 0; i < repeatFactor; i++ {
		repeatedString += character
	}
	return repeatedString
}
