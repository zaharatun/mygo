package handlers

import (
	"net/http"
	


)

func Home(w http.ResponseWriter, r *http.Request) {
	RenderT(w,"home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTapmlate(w,"about.page.tmpl")
}