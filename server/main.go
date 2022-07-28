package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "it is response mate welcome!")
	fmt.Println("can YOu see this?")
}

func hola(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to hola page!")
}

func handler() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/hola", hola)
}

func main() {
	handler()
	log.Fatal(http.ListenAndServe(":8080", nil))

}
