package api

import "net/http"

func checkHealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ok"))
}
