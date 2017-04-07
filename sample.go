package sample

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", index)
	http.HandleFunc("/nice", nice)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World 2")
}

func nice(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "very nice")
}
