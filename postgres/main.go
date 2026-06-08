package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	_ "github.com/lib/pq"
)

var db *sql.DB
var tpl *template.Template

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://bookworm:password@localhost:5432/bookstore?sslmode=disable")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

type Books struct {
	Isbn   string
	Title  string
	Author string
	Price  float32
}

func main() {
	http.HandleFunc("/books", getBooks)
	http.HandleFunc("/book", getBook)
	http.HandleFunc("/book/create", createBookForm)
	http.HandleFunc("/book/create/book", createBook)
	http.HandleFunc("/book/update", updateBookForm)
	http.HandleFunc("/book/update/book", updateBook)
	http.HandleFunc("/book/delete", deletebook)
	http.ListenAndServe(":8080", nil)
}

func getBooks(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	rows, err := db.Query("Select * from books;")
	if err != nil {
		log.Printf("Query error: %v", err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	books := make([]Books, 0)
	for rows.Next() {
		book := Books{}
		err := rows.Scan(&book.Isbn, &book.Title, &book.Author, &book.Price)
		if err != nil {
			log.Printf("Scan error: %v", err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		books = append(books, book)
	}
	if err = rows.Err(); err != nil {
		log.Printf("Rows iteration error: %v", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	err = tpl.ExecuteTemplate(w, "books.gohtml", books)
	if err != nil {
		log.Printf("Template error: %v", err)
		http.Error(w, http.StatusText(500), 500)
	}
}
func getBook(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	isbn := req.FormValue("isbn")
	if isbn == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}
	row := db.QueryRow("Select * from books where isbn=$1", isbn)
	book := Books{}
	err := row.Scan(&book.Isbn, &book.Title, &book.Author, &book.Price)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, req)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), 500)
		return
	}
	err = tpl.ExecuteTemplate(w, "book-details.gohtml", book)
	if err != nil {
		log.Printf("Template error: %v", err)
		http.Error(w, http.StatusText(500), 500)
	}
}
func createBookForm(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "insert.gohtml", nil)
}
func createBook(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	book := Books{}
	book.Isbn = req.FormValue("isbn")
	book.Title = req.FormValue("title")
	book.Author = req.FormValue("author")
	price := req.FormValue("price")

	if book.Isbn == "" && book.Title == "" && book.Author == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}
	f32, err := strconv.ParseFloat(price, 32)
	if err != nil {
		http.Error(w, http.StatusText(406), 406)
		return
	}
	book.Price = float32(f32)
	_, err = db.Exec("Insert into books(isbn,title,author,price) values($1,$2,$3,$4)", book.Isbn, book.Title, book.Author, book.Price)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	tpl.ExecuteTemplate(w, "insert.gohtml", book)
}
func updateBookForm(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	isbn := req.FormValue("isbn")
	if isbn == "" {
		http.Error(w, http.StatusText(400)+" - ISBN required", 400)
		return
	}
	row := db.QueryRow("SELECT * FROM books WHERE isbn=$1", isbn)
	book := Books{}
	err := row.Scan(&book.Isbn, &book.Title, &book.Author, &book.Price)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, req)
		return
	case err != nil:
		log.Printf("Scan error: %v", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	err = tpl.ExecuteTemplate(w, "update.gohtml", book)
	if err != nil {
		log.Printf("Template error: %v", err)
		http.Error(w, http.StatusText(500), 500)
	}
}
func updateBook(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	book := Books{}
	book.Isbn = req.FormValue("isbn")
	book.Title = req.FormValue("title")
	book.Author = req.FormValue("author")
	price := req.FormValue("price")

	if book.Isbn == "" || book.Title == "" || book.Author == "" {
		http.Error(w, http.StatusText(400)+" - All fields required", 400)
		return
	}

	f64, err := strconv.ParseFloat(price, 64)
	if err != nil {
		http.Error(w, http.StatusText(406)+" - Invalid price", 406)
		return
	}
	book.Price = float32(f64)

	_, err = db.Exec("UPDATE books SET title=$1, author=$2, price=$3 WHERE isbn=$4",
		book.Title, book.Author, book.Price, book.Isbn)
	if err != nil {
		log.Printf("Update error: %v", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	http.Redirect(w, req, "/book?isbn="+book.Isbn, http.StatusSeeOther)
}
func deletebook(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	isbn := req.FormValue("isbn")
	if isbn == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	_, err := db.Exec("Delete from books Where isbn=$1;", isbn)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	http.Redirect(w, req, "/books", http.StatusSeeOther)
}
