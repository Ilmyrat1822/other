package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", index)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/remove", remover)
	http.ListenAndServe(":7777", nil)

}

func index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, `<a href="/set"  >Set cookies</a>`)
}
func set(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: "mycookies", Value: "value"})
	fmt.Fprintln(w, `<a href="/read"  >Read cookies</a>`)
}

func read(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("mycookies")
	if err != nil {
		log.Printf("Cookie error: %v", err)
		http.Redirect(w, r, "/set", http.StatusSeeOther)
		return
	}

	fmt.Fprintln(w, `<a href="/remove">Remove cookies</a>`)
	fmt.Fprintln(w, `<h2>`+html.EscapeString(c.Name)+`</h2>`)
}

func remover(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("mycookies")

	if err != nil {
		log.Printf("Cookie error: %v", err)
		http.Redirect(w, r, "/set", http.StatusSeeOther)
	}
	c.MaxAge = -1
	http.SetCookie(w, c)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
