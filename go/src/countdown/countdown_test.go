package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {

	t.Run("testing Countdown output and number of Sleep calls", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}
		Countdown(buffer, spySleeper)
		actualOutput := buffer.String()
		expectedOutput := "3\n2\n1\nGo!"
		if actualOutput != expectedOutput {
			t.Errorf("Actual output was %q but expected %q", actualOutput, expectedOutput)
		}

		if spySleeper.Calls != 4 {
			t.Errorf("Expected 4 calls to spySleeper but only saw %d", spySleeper.Calls)
		}
	})

	t.Run("testing Countdown operations order", func(t *testing.T) {
		operationsTracker := &CountdownOperationsTracker{}
		Countdown(operationsTracker, operationsTracker)

		expectedQueue := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(expectedQueue, operationsTracker.Operations) {
			t.Errorf("Expected calls %v but got %v", expectedQueue, operationsTracker.Operations)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("Expected to sleep %v seconds but slept %v seconds", sleepTime, spyTime.durationSlept)
	}
}
