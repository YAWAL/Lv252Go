package main

import (
	"net/http"
	"fmt"
	"log"
)

func main() {
	http.HandleFunc("/", handler)//Каждый запрос вызывает обработчик
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
// Обработчик возвращает компонент пути из URL запроса,
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{\n \"name\":%q \n}", r.URL.Path) //URL.Path =
}
