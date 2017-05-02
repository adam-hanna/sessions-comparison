package benchmark

import (
	"net/http"
	"testing"
)

func BenchmarkValidSession(b *testing.B) {
	// now let's get a valid session
	res, err := http.Get("http://localhost:8080/issue")
	if err != nil {
		b.Errorf("Couldn't send request to test server; Err: %v\n", err)
	}

	if res.StatusCode != 200 {
		b.Fatalf("Expected (200), received: %d\n", res.StatusCode)
	}

	// now let's send to the require server
	// first, grab the session cookie
	rc := res.Cookies()
	var sessionCookieIndex int
	for i, cookie := range rc {
		if cookie.Name == "session-key" {
			sessionCookieIndex = i
		}
	}

	req, err := http.NewRequest("GET", "http://localhost:8080/require", nil)
	if err != nil {
		b.Fatalf("Couldn't build request; Err: %v\n", err)
	}

	req.AddCookie(rc[sessionCookieIndex])

	tr := &http.Transport{}
	defer tr.CloseIdleConnections()
	cl := &http.Client{
		Transport: tr,
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, err := cl.Do(req)
		if err != nil {
			b.Fatal("Get:", err)
		}
		if res.StatusCode != 200 {
			b.Fatalf("Wanted 200 status code, received: %d\n", res.StatusCode)
		}
	}
}
