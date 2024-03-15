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
	http.HandleFunc("/", helloHandler)

	log.Println("sever start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
