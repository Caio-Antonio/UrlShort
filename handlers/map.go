package handlers

import (
	"net/http"
)

func MapHandler(pathToURls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request){
		path := r.URL.Path
		if dest, ok := pathToURls[path]; ok {
			http.Redirect(rw, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(rw, r)
		} 
 }