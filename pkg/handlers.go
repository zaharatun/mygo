package main

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renderTapmlate(w,"home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTapmlate(w,"about.page.tmpl")
}