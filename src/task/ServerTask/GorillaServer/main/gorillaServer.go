package main

import (
	"github.com/gorilla/mux"
	"net/http"

	"log"
)

type myName string

func main() {
	rtr := mux.NewRouter()
	//rtr.HandleFunc("/", pageHandler)
	http.Handle("/", rtr)
	//http.HandleFunc("/", handler)//
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//func handler(w http.ResponseWriter, r *http.Request)  {
//	log.Fatal(w, "", r.URL.p)
//}

//func handler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "{\n \"name\":%q \n}", r.URL.Path)
//}