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
		fmt.Fprintf(w, "Hello, world!\n")
		for _, env := range os.Environ() {
			fmt.Fprintf(w, "%s\n", env)
		}
	})

	message("listening on port %s\n", port)

	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
