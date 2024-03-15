package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!")
	}

	postArticleHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "post article...\n")
	}

	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/article", postArticleHandler)
	log.Println("sever start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
