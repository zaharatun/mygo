package render

import (
	"fmt"
	"net/http"
	"html/template"
)


func RenderTapmlate(w http.ResponseWriter, tmpl string) {
prasedTamplate, _ := template.ParseFiles("./templates/" + tmpl)
err := prasedTamplate.Execute(w, nil)
 if err != nil {
	fmt.Println("error parsing template:", err)
	return
}

}
