package arrays

/*ArraySum ... Function Specification:
INPUTS:
	numbers = An array of integer values
OUTPUTS:
	The sum of all the entries in the inputted array
*/
func ArraySum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

/*ArraySumAll ... Function Specification:
INPUTS:
	numbersToSum = A variable length array of arrays of integers
OUTPUTS:
	sums = An array with the sums of the inputted arrays
*/
func ArraySumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, ArraySum(numbers))
	}
	return
}
