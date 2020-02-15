package main

import (
	"fmt"
	"io"
	"os"
)

/*Greet ...
INPUTS:
	writer = Input io.writer object to write to
	name = String to write to the inputted buffer
OUTPUTS:
	None, writes the inputted name string to the
	inputted bytes Buffer
*/
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	Greet(os.Stdout, "Shantanu")
}
