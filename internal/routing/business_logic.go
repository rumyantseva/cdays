package routing

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewBLRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/home", rootHandler())
	return r
}

func rootHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello!"))
	}
}
