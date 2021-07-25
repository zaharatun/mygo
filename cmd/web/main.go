package main


import(
	"fmt"
	"net/http"
	"pkg/handlers"
)


const portNuber =":2030"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Println("Starting application on port:",portNumber)
	http.ListenAndServe(portNumber, nil)
}