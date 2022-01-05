package httpHandler

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Users")
}

func NewHttpHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", IndexHandler)
	mux.HandleFunc("/user", GetUsersHandler)
	return mux
}
