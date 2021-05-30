package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	_ = http.ListenAndServe("0.0.0.0:61000", nil)
}
