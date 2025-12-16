package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	Handlefunc()
}

type Articles struct {
	Id                    int
	Title, Desc, Fulltext string
}

var posts = []Articles{}
var showpost = Articles{}

func Handlefunc() {
	router := mux.NewRouter()
	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/create/", create).Methods("GET")
	router.HandleFunc("/save_article/", save_article).Methods("POST")
	router.HandleFunc("/post/{id:[0-9]+}", post).Methods("GET")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("wwwroot/static/"))))

	http.ListenAndServe(":7777", router)
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("wwwroot/index.html", "wwwroot/header.html", "wwwroot/footer.html")
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}

	db, err := sql.Open("mysql", "root:sanly2024@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	res, err := db.Query("SELECT * FROM articles")
	if err != nil {
		panic(err)
	}
	defer res.Close()

	posts = []Articles{}
	for res.Next() {
		var post Articles
		err = res.Scan(&post.Id, &post.Title, &post.Desc, &post.Fulltext)
		if err != nil {
			panic(err)
		}

		posts = append(posts, post)
	}

	t.ExecuteTemplate(w, "index", posts)

}
func create(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("wwwroot/create.html", "wwwroot/header.html", "wwwroot/footer.html")
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}

	t.ExecuteTemplate(w, "create", nil)

}
func save_article(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	desc := r.FormValue("desc")
	full_text := r.FormValue("full_text")

	if title == "" || desc == "" || full_text == "" {
		fmt.Fprintf(w, "Fill all blanks")
	} else {

		db, err := sql.Open("mysql", "root:sanly2024@tcp(127.0.0.1:3306)/golang")
		if err != nil {
			panic(err)
		}

		defer db.Close()

		insert, err := db.Query(fmt.Sprintf("INSERT INTO articles (`title`,`desc`,`fulltext`) Values('%s','%s','%s')", title, desc, full_text))
		if err != nil {
			panic(err)
		}

		defer insert.Close()

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
func post(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	t, err := template.ParseFiles("wwwroot/post.html", "wwwroot/header.html", "wwwroot/footer.html")
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}

	db, err := sql.Open("mysql", "root:sanly2024@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	res, err := db.Query(fmt.Sprintf("SELECT * FROM articles WHERE id='%s'", vars["id"]))
	if err != nil {
		panic(err)
	}
	defer res.Close()

	showpost = Articles{}
	for res.Next() {
		var post Articles
		err = res.Scan(&post.Id, &post.Title, &post.Desc, &post.Fulltext)
		if err != nil {
			panic(err)
		}

		showpost = post
	}

	t.ExecuteTemplate(w, "post", showpost)
}
