package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/pat"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello World!")
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	user := User{
		Name:  "kelly",
		Email: "kelly@gmail.com",
		Age:   33,
	}
	data, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}

func NewHttpHandler() http.Handler {
	mux := pat.New()
	// mux.HandleFunc("/", indexHandler).Methods("GET")
	// mux.HandleFunc("/users", getUsersHandler).Methods("GET")
	mux.Get("/", indexHandler)
	mux.Get("users", getUsersHandler)
	return mux
}
