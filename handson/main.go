package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var User = map[string]string{}
var tmpl = template.Must(template.ParseGlob("static/*.html"))
var key = []byte("mysecretkey")

func main() {
	fmt.Println("Server successfully started")
	http.HandleFunc("/", index)
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.ListenAndServe(":8080", nil)
}
func index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Redirect(w, r, "Not Allowed", http.StatusSeeOther)
		return
	}
	tmpl.ExecuteTemplate(w, "index.html", nil)
}
func register(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "Not Allowed", http.StatusSeeOther)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	if username == "" || password == "" {
		http.Redirect(w, r, "Credentials can not be empty", http.StatusSeeOther)
		return
	}
	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Redirect(w, r, "Password hashing error", http.StatusInternalServerError)
		return
	}
	User[username] = string(hashedpassword)
	fmt.Println("Registered users:", User)
	tmpl.ExecuteTemplate(w, "login.html", nil)
}
func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl.Execute(w, nil)
		return
	}
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "" || password == "" {
			tmpl.Execute(w, map[string]string{"error": "Credentials cannot be empty"})
			return
		}

		hashedPassword, exists := User[username]
		if !exists {
			tmpl.Execute(w, map[string]string{"error": "User not found"})
			return
		}

		err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
		if err != nil {
			tmpl.Execute(w, map[string]string{"error": "Invalid password"})
			return
		}

		tmpl.ExecuteTemplate(w, "dashboard.html", username)
	}
}
func createToken(sid string) string {

	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(sid))
	sigendMac := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return sigendMac + "|" + sid
}
func parseToken(ss string) (string, error) {
	xs := strings.SplitN(ss, "|", 2)
	if len(xs) != 2 {
		return "", fmt.Errorf("Not enough string")
	}
	b64 := xs[0]
	xb, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return "", fmt.Errorf(" Could not parse token %w", err)
	}
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(xs[1]))

	ok := hmac.Equal(xb, mac.Sum(nil))
	if !ok {
		return "", fmt.Errorf("Could not parse token")
	}
	return xs[1], nil
}
