package main

import (
	"fmt"
	"net/http"
	"time"
)

type User struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Age       int       `json:"age_name"`
	CreateAt  time.Time `json:"create_at"`
}

// struct
type Foo struct {
	value string
}

func (f *Foo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello "+f.value+"!")
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Bar!")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	http.HandleFunc("/bar", bar)

	http.Handle("/foo", &Foo{value: "foo"})

	http.ListenAndServe(":3000", nil)
}
