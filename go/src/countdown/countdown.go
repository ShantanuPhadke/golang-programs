package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// SLICING UP REQUIREMENTS:
//	[*] (1) Get the program to print just '3'
//	[*] (2) Get the program to print 3,2,1 and Go! on separate lines
//  [*] (3) Get the program to wait one second between each line

/*Sleeper ... interface that requires Sleep method implementation*/
type Sleeper interface {
	Sleep()
}

/*SpySleeper struct that mimics Sleep calls via the Calls integer attribute*/
type SpySleeper struct {
	Calls int
}

/*Sleep ... our "mocked" sleeping function for our tests to call*/
func (s *SpySleeper) Sleep() {
	s.Calls++
}

/*RealSleeper struct we will use for our actual calls to Countdown*/
type RealSleeper struct{}

/*Sleep ... sleeper function that will be used by our actual program's calls to Countdown*/
func (rs *RealSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

/*CountdownOperationsTracker ... A struct that keeps track of operation order*/
type CountdownOperationsTracker struct {
	Operations []string
}

/*write constant to keep track of write countdown operations*/
const write = "write"

/*sleep constant to keep track of sleep countdown operations*/
const sleep = "sleep"

/*Sleep ... CountdownOperationsTracker for writing a sleep operation to our queue*/
func (s *CountdownOperationsTracker) Sleep() {
	s.Operations = append(s.Operations, sleep)
}

/*Write ... CountdownOperationsTracker for writing a write operation to our queue*/
func (s *CountdownOperationsTracker) Write(buffer []byte) (n int, err error) {
	s.Operations = append(s.Operations, write)
	return
}

/*ConfigurableSleeper ... Literally a struct to represent Sleepers with customizable sleep times*/
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

/*Sleep ... Implementing the Sleep function for the ConfigurableSleeper type*/
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

/*SpyTime ... struct with durationSlept attribute*/
type SpyTime struct {
	durationSlept time.Duration
}

/*Sleep ...
INPUTS:
	duration = how much time we want to sleep
*/
func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func main() {
	realCustomizedSleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, realCustomizedSleeper)
}

const finalWord = "Go!"
const countdownStart = 3

/*Countdown ...
INPUTS:
	writer = The buffer we want to write our output(s) to
*/
func Countdown(writer io.Writer, s Sleeper) {
	for outputCount := countdownStart; outputCount > 0; outputCount-- {
		s.Sleep()
		fmt.Fprintln(writer, outputCount)
	}
	s.Sleep()
	fmt.Fprint(writer, finalWord)
}
