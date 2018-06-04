package main

import (
	"io"
	"net/http"
)

func copyResponse(w http.ResponseWriter, r *http.Response) {
	for k, v := range r.Header {
		w.Header().Set(k, v[0])
	}
	w.WriteHeader(r.StatusCode)
	io.Copy(w, r.Body)
	r.Body.Close()
}
