package handlers

import (
	"net/http"
	"github.com/zaharatun/mygo/pkg/render"

)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTapmlate(w,"home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTapmlate(w,"about.page.tmpl")
}