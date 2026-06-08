package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("wwwroot/html/*"))
}

func main() {
	handle()
}

func handle() {
	http.Handle("/wwwroot/", http.StripPrefix("/wwwroot/", http.FileServer(http.Dir("wwwroot/"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", index)
	addr := fmt.Sprintf("0.0.0.0:%v", 7777)
	fmt.Println("Starting server on " + addr)
	http.ListenAndServe(":7777", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index", nil)
}
