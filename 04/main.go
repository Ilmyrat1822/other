package main

import (
	"fmt"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func main() {
	handle()
}

func handle() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", index)
	http.ListenAndServe(":7777", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	var cookie *http.Cookie
	_, err := r.Cookie("session")
	if err != nil {
		log.Println(err)
		i := uuid.NewV4()
		cookie = &http.Cookie{Name: "session",
			Value: i.String(),
			//Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)

		fmt.Println(cookie.Value)
	}

}
