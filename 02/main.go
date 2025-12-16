package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/file/", filef)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	//http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("wwwroot/assets"))))
	http.ListenAndServe(":7777", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("q")
	io.WriteString(w, "Do my search: "+v)
}

func filef(w http.ResponseWriter, r *http.Request) {

	f, h, err := r.FormFile("q")

	if err != nil {
		log.Fatal(err)
	}

	bs, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(bs)
	}
	fmt.Println(h.Filename)

	s := string(bs)
	fmt.Println(s)
}
