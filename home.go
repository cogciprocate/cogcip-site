package main

import (
	"net/http"
	//"github.com/nsan1129/unframed/log"
	//"github.com/nsan1129/unframed"
)

func homeReg() {
	net.Get("/", home)
	// net.Get("/help", cc_help)

	net.TemplateFiles(
		"tmpl/home.html.tmpl",
	)
}

func home(w http.ResponseWriter, r *http.Request) {
	net.ExeTmpl(w, "home", nil)
}

// func cc_help(w http.ResponseWriter, r *http.Request) {
// 	net.ExeTmpl(w, "help_cc", nil)
// }
