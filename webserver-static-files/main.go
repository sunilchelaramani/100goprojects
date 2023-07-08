//simple web server that serves static files

package main

import (
	"fmt"
	"net/http"
)

func main() {
	//handle static files
	http.Handle("/", http.FileServer(http.Dir("./static")))
	//handle dynamic files
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
	//start server
	http.ListenAndServe(":8080", nil)
}
