package main

import (
	"fmt"
	"net/http"
	"os"
)

var port string = "3000"

func message(format string, a ...interface{}) {
	fmt.Printf(fmt.Sprintf("tug-example: %s", format), a...)
}

func main() {
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	})

	message("listening on port %s\n", port)

	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
