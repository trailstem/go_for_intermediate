package main

import (
	"fmt"
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
	postListHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "postListHandler...\n")
	}
	postDetailHandler := func(w http.ResponseWriter, req *http.Request) {
		articleID := 1
		resString := fmt.Sprintf("article_%d", articleID)
		io.WriteString(w, resString)

	}
	postNiceHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Nice...\n")
	}
	postCommentHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Comment...\n")
	}

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/article", postArticleHandler)
	http.HandleFunc("/article/list", postListHandler)
	http.HandleFunc("/article/1", postDetailHandler)
	http.HandleFunc("/article/nice", postNiceHandler)
	http.HandleFunc("/comment", postCommentHandler)
	log.Println("sever start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
