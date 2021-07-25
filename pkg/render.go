package main

import (
	"fmt"
	"net/http"
	"text/template"
)


func renderTapmlate(w http.ResponseWriter, tmpl string) {
prasedTamplate,_ :=template.ParseFiles("./templates/" +tmpl)
err := prasedTamplate.Execute(w, nil)
 if err != nil {
	fmt.Println("error parsing template:", err)
}

}
