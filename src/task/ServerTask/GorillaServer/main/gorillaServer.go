package main

import (
	"github.com/gorilla/mux"
	"net/http"

	"log"
)

func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/{name:[a-z]+}", profile).Methods("GET")
	http.Handle("/", rtr)
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}

func profile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]
	w.Write([]byte("{\n \"name\":\"" + name + "\"\n}"))
}

