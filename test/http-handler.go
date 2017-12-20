package main

import (
	"fmt"
	"log"
	"net/http"
)

type String string

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.Response.Body)
}
func (st Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.Response)
}

func main() {
	// your http.Handle calls here
	http.Handle("/string", String("I'm frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})

	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}

//type Hello struct{}

//func (h Hello) ServeHTTP(	w http.ResponseWriter,r *http.Request) {
//	fmt.Fprint(w, "Hello!")
//}

//func main() {
//	var h Hello
//	err := http.ListenAndServe("localhost:4000", h)
//	if err != nil {
//		log.Fatal(err)
//	}
//}
//http.Handle("/string", String("I'm a frayed knot."))
//http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
