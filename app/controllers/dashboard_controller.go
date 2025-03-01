package controllers

import (
	"net/http"
	"time"

	"github.com/unrolled/render"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {

	render := render.New(render.Options{
		Layout: "layout",
	})

	timeFormated := time.Now().Format("02 February, 2006")

	_ = render.HTML(w, http.StatusOK, "dashboard", map[string]interface{}{
		"title":    "Dashboard",
		"username": "Asep",
		"date":     timeFormated,
		"body":     "Description",
	})
}
