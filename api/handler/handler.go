package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Get users")
}

func handleGetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, "User ID : ", vars["id"])
}

func NewHttpHandler() http.Handler {
	mux := mux.NewRouter()
	// mux := http.NewServeMux()
	mux.HandleFunc("/", handleIndex)
	mux.HandleFunc("/users", handleUsers)
	mux.HandleFunc("/users/{id:[0-9]+}", handleGetUser)
	return mux
}
