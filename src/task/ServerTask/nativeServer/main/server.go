package main

import (
	"net/http"
	"fmt"
	"log"
)

func main() {
	http.HandleFunc("/", handler)//
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
//return text after / in URL Path
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{\n \"name\":%q \n}", r.URL.Path)
}
