package main

import (
	"log"
	"net/http"

	"github.com/trailstem/go_for_intermediate/handlers"
)

func main() {

	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/article", handlers.PostArticleHandler)
	http.HandleFunc("/article/list", handlers.PostListHandler)
	http.HandleFunc("/article/1", handlers.PostDetailHandler)
	http.HandleFunc("/article/nice", handlers.PostNiceHandler)
	http.HandleFunc("/comment", handlers.PostCommentHandler)
	log.Println("sever start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
