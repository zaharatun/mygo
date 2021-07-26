package main


import(
	"fmt"
	"net/http"
	"gitub.com/zaharatun/mygo/pkg/handlers"
	
)


const portNumber =":2030"

func main() {
	http.HandleFunc("/", handalers.Home)
	http.HandleFunc("/about", handalers.About)
	fmt.Println("Starting application on port:",portNumber)
	http.ListenAndServe(portNumber, nil)
}