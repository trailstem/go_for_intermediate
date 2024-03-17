package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/trailstem/go_for_intermediate/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", handlers.PostListHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/1", handlers.PostDetailHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/nice", handlers.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", handlers.PostCommentHandler).Methods(http.MethodPost)
	log.Println("sever start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
