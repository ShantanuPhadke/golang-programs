package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	// Mocking both the slow and fast website servers
	// These calls actually open up HTTP servers on any of the open ports
	slowServer := makeMockServer(20 * time.Millisecond)
	fastServer := makeMockServer(0 * time.Millisecond)

	// The defer keyword will execute functions / lines of code at the end of
	// the function that contains the defered lines of code. In this case we
	// are closing our mock servers after testing them
	defer slowServer.Close()
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	expectedURL := fastURL
	actualURL := Racer(slowURL, fastURL)

	if expectedURL != actualURL {
		t.Errorf("Expected %q, wanted %q", expectedURL, actualURL)
	}
}

func TestConcurrentRacer(t *testing.T) {
	t.Run("testing basic concurrent racing with one slow and one fast server", func(t *testing.T) {
		t.Helper()
		// Mocking both the slow and fast website servers
		// These calls actually open up HTTP servers on any of the open ports
		slowServer := makeMockServer(20 * time.Millisecond)
		fastServer := makeMockServer(0 * time.Millisecond)

		// The defer keyword will execute functions / lines of code at the end of
		// the function that contains the defered lines of code. In this case we
		// are closing our mock servers after testing them
		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		expectedURL := fastURL
		actualURL, _ := ConcurrentRacer(slowURL, fastURL, 10*time.Millisecond)

		if expectedURL != actualURL {
			t.Errorf("Expected %q, wanted %q", expectedURL, actualURL)
		}
	})

	t.Run("testing that the ConcurrentTester returns an error when both urls take longer than 10s to return", func(t *testing.T) {
		t.Helper()
		slowServer1 := makeMockServer(11 * time.Millisecond)
		slowServer2 := makeMockServer(12 * time.Millisecond)
		defer slowServer1.Close()
		defer slowServer2.Close()
		slowURL1 := slowServer1.URL
		slowURL2 := slowServer2.URL
		_, err := ConcurrentRacer(slowURL1, slowURL2, 10*time.Millisecond)
		if err == nil {
			t.Error("Expected an error but got none")
		}
	})
}

func makeMockServer(timeDelay time.Duration) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(timeDelay)
		w.WriteHeader(http.StatusOK)
	}))

	return server
}
