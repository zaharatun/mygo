package main

import (
	"net/http"
	"github.com/zaharatun/mygo/pkg/handlers"
)




func main() {
	http.HandleFunc("/",handlers.Home)
	http.HandleFunc("/",handlers.About)

	http.ListenAndServe(":2030", nil)
}