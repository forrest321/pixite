package handlers

import (
	"github.com/forrest321/pixite/constants"
	"github.com/forrest321/pixite/models"
	"github.com/gorilla/context"
	"html/template"
	"net/http"
)

var imageTmpl = template.Must(template.ParseFiles("templates/base.html.tmpl", "templates/image.html.tmpl"))

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	serveable, ok := context.GetOk(r, constants.ServeableOneKey).(models.Serveable)
	if !ok {
		imageTmpl.Execute(w, nil)
	} else {
		imageTmpl.Execute(w, serveable)
	}
}
