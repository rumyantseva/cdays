package routing_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rumyantseva/cdays/internal/routing"
)

func TestNewBLRouter(t *testing.T) {
	r := routing.NewBLRouter()
	srv := httptest.NewServer(r)
	defer srv.Close()

	testCases := []struct {
		route              string
		expectedStatusCode int
	}{
		{"/home", http.StatusOK},
		{"/", http.StatusNotFound},
		// ...
		// ...
	}

	for _, c := range testCases {
		resp, err := http.Get(srv.URL + c.route)
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != c.expectedStatusCode {
			t.Errorf("Status code is %v, but %v expected", resp.StatusCode, c.expectedStatusCode)
		}
	}
}
