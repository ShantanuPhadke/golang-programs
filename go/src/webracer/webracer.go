package racer

import (
	"fmt"
	"net/http"
	"time"
)

/*Racer ... function that "races" two urls against each other
INPUTS:
	slowURL - The string of the URL we expect to be slower
	fastURL - The string of the URL we expect to be faster
OUTPUTS:
	Returns string of the faster URL
*/
func Racer(slowURL, fastURL string) (raceWinner string) {
	slowURLDuration := measureResponseTime(slowURL)
	fastURLDuration := measureResponseTime(fastURL)

	if slowURLDuration < fastURLDuration {
		return slowURL
	}

	return fastURL
}

func measureResponseTime(url string) time.Duration {
	startTime := time.Now()
	http.Get(url)
	URLDuration := time.Since(startTime)
	return URLDuration
}

/*ConcurrentRacer ... A concurrent version of our Racer from above
INPUTS:
	slowURL - The string of the URL we expect to be slower
	fastURL - The string of the URL we expect to be faster
OUTPUTS:
	Returns string of the faster URL
*/
func ConcurrentRacer(slowURL, fastURL string, timeout time.Duration) (raceWinner string, err error) {
	select {
	case <-ping(slowURL):
		return slowURL, nil
	case <-ping(fastURL):
		return fastURL, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for slow urls %s and %s", slowURL, fastURL)
	}
}

/*ping ...
INPUTS:
	url - The string of the url we are trying to ping
OUTPUTS:
	Returns a channel with a struct
*/
func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
