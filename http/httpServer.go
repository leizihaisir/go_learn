package main

import (
	"net/http"
	"log"
)

func sayBay(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bye bye,this is version 1"))
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello,this is version 1"))
	})
	http.HandleFunc("/bye", sayBay)
	log.Println("starting server ... v1")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
