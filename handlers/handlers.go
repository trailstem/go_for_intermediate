package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!")
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "post article...\n")
}
func PostListHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "postListHandler...\n")
}
func PostDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID := 1
	resString := fmt.Sprintf("article_%d", articleID)
	io.WriteString(w, resString)

}

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID := 1
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}
	resString := fmt.Sprintf("Article No.%d\n", articleID)
	io.WriteString(w, resString)
}

func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Nice...\n")
}
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Comment...\n")
}
