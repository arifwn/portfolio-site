package portfolio

import (
	"html/template"
	"net/http"
)

type Page struct {
	Title string
}

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	page := Page{Title: "Arif Widi Nugroho"}
	t, _ := template.ParseFiles("/home/arif/projects/sainsmograf/templates/index.html")
	t.Execute(w, page)
}
