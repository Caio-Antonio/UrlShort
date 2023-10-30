package main

import (
	"net/http"

	"github.com/Caio-Antonio/UrlShort/handlers"
)

func defaultMux() *http.ServeMux{
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HelloHandler)
	return mux
}

func main (){
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := handlers.MapHandler(pathsToUrls, mux)
	http.ListenAndServe(":8080", mapHandler)
}