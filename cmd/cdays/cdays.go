package main

import "net/http"

func main() {
	http.HandleFunc("/", rootHandler())
	http.ListenAndServe(":8000", nil)
}

func rootHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello!"))
	}
}
