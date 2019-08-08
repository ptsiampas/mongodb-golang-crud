package main

import (
	"log"
	"mongodb-web-dev/books"
	"net/http"
)

func main() {

	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("./css"))))
	http.HandleFunc("/", books.Index)
	http.HandleFunc("/book/read", books.GetOneBook)
	http.HandleFunc("/book/create", books.StoreOneBook)
	http.HandleFunc("/book/update", books.UpdateOneBook)
	http.HandleFunc("/book/delete", books.RemoveOneBook)

	log.Print("Listening on http://localhost:8008")
	log.Fatalln(http.ListenAndServe("0.0.0.0:8080", nil))
}
