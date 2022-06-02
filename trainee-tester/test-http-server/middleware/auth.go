package middleware

import "net/http"

func AuthName(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name != "Sergey" {
			http.Error(w, "Wrong name", http.StatusUnauthorized)
			return
		}
		h(w, r)
	}
}
