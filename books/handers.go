package books

import (
	"github.com/ptsiampas/mongodb-golang-crud/config"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"

	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	isGet(w, r)

	bks := AllBooks()

	err := config.TPL.ExecuteTemplate(w, "index.gohtml", bks)
	if err != nil {
		log.Panicln(err)
	}
}

func GetOneBook(w http.ResponseWriter, r *http.Request) {
	isGet(w, r)

	isbn := r.FormValue("isbn")

	if isbn == "" {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		return
	}

	bk, err := FindOneBook(isbn)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if err := config.TPL.ExecuteTemplate(w, "book.gohtml", bk); err != nil {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
	}

}

func UpdateOneBook(w http.ResponseWriter, r *http.Request) {

	// TODO: This does not abstract the Model cleanly; Since I need to know about the id type, Should Fix
	if r.Method == "POST" {
		bid, _ := primitive.ObjectIDFromHex(r.FormValue("s"))
		b := Book{
			Id:     bid,
			Isbn:   r.FormValue("isbn"),
			Title:  r.FormValue("title"),
			Author: r.FormValue("author"),
		}
		p := r.FormValue("price")
		b.Price, _ = strconv.ParseFloat(p, 64)

		oid, err := UpdateBook(b)
		if err != nil {
			log.Println("UpdateOneBook", err)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		http.Redirect(w, r, "/book/get?isbn="+oid, http.StatusSeeOther)
		return
	}

	// Make sure its a GET method.
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	isbn := r.FormValue("isbn")

	if isbn == "" {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
	}

	// Get the book from the database.
	b, err := FindOneBook(isbn)
	if err != nil {
		log.Println("Unable to find a book:", err)
		return
	}

	if err := config.TPL.ExecuteTemplate(w, "update.gohtml", b); err != nil {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
	}

}

func StoreOneBook(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" { // Collect all the details
		b := Book{
			Isbn:   r.FormValue("isbn"),
			Title:  r.FormValue("title"),
			Author: r.FormValue("author"),
		}
		p := r.FormValue("price")
		b.Price, _ = strconv.ParseFloat(p, 64)

		oid, err := AddBook(b)
		if err != nil {
			log.Println("store-one-book", err)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		http.Redirect(w, r, "/book/get?isbn="+oid, http.StatusSeeOther)
	}

	if err := config.TPL.ExecuteTemplate(w, "insert.gohtml", nil); err != nil {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
	}

}

func RemoveOneBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" || r.FormValue("isbn") == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	_, err := DeleteBook(r.FormValue("isbn"))
	if err != nil {
		log.Println(err)
		return
	}

	// TODO: Should probably replace this with a page.
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func isGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
}
