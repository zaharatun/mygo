package main

import (
	"github.com/zaharatun/mygo/helpers"
	"log"
)

func main() {
	log.Println("hello")
	var myVar helpers.SomeType
	myVar.TypeName = "zahir"
	myVar.TypeNmber = 10
	log.Println(myVar.TypeName)
	log.Println(myVar.TypeNumber)
}
