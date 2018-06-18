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

	{
		resp, err := http.Get(srv.URL + "/home")
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Status code is %v, but %v expected", resp.StatusCode, http.StatusOK)
		}
	}

	{
		resp, err := http.Get(srv.URL)
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("Status code is %v, but %v expected", resp.StatusCode, http.StatusNotFound)
		}
	}

}
