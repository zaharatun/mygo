package main

import (
	"github.com/zaharatun/mygo/helpers"
	"log"
)

func main() {
	log.Println("hello")
	var myVar helpers.SomeType
	myVar.TypeName = "zahir"
	log.Println(myVar.TypeName)
}
