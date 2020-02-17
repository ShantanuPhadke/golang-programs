package counter

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("increment counter 3 times, leave it as 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()
		assertCounterValue(t, counter, 3)
	})

	t.Run("test the counter via concurrent calls", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			// Concurrent go routine
			go func(w *sync.WaitGroup) {
				counter.Inc()
				w.Done()
			}(&wg)
		}

		wg.Wait()

		assertCounterValue(t, counter, wantedCount)
	})
}

func assertCounterValue(t *testing.T, currentCounter *Counter, expectedCount int) {
	t.Helper()
	if currentCounter.Value() != expectedCount {
		t.Errorf("Expected %d but got %d", expectedCount, currentCounter.Value())
	}
}
