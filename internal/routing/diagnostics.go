package routing

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func NewDiagnosticsRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/healthz", handleOK())
	r.HandleFunc("/readyz", handleOK())
	return r
}

func handleOK() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, http.StatusText(http.StatusOK))
	}
}
