package handlers

import "net/http"

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write("stuff to write")

}
