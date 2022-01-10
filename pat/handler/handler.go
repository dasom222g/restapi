package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/pat"
	"github.com/unrolled/render"
)

var rd *render.Render

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
		Age:   25,
	}

	/*
		data, _ := json.Marshal(user)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(data))
	*/
	// render로 윗 부분 소스 축약
	rd.JSON(w, http.StatusOK, user)
}

func NewHttpHandler() http.Handler {
	rd = render.New() // rd 정의

	mux := pat.New()
	mux.Get("/users", getUsersHandler)
	mux.Get("/", indexHandler)
	return mux
}
