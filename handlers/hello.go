package handlers

import (
	"fmt"
	"net/http"
)


 func HelloHandler(rw http.ResponseWriter, r *http.Request){
	fmt.Printf("Hello World")
 }