package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gorilla/pat"
	"github.com/unrolled/render"
)

var rd *render.Render
var currentID int
var users []*User

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello World!")
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	/*
		data, _ := json.Marshal(users)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(data))
	*/
	// render로 윗 부분 소스 축약
	rd.JSON(w, http.StatusOK, users)
}

func addUsersHandler(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		// w.WriteHeader(http.StatusBadRequest)
		// fmt.Fprint(w, err)
		rd.Text(w, http.StatusBadRequest, err.Error())
		return
	}
	currentID++
	user.ID = currentID
	user.CreatedAt = time.Now()

	users = append(users, user)

	/*
		data, err := json.Marshal(user)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(data))
	*/

	rd.JSON(w, http.StatusOK, user)

}

func templateHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("tmpl").ParseFiles("templates/hello.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	tmpl.ExecuteTemplate(w, "hello.html", "dasomi")
}

func NewHttpHandler() http.Handler {
	currentID = 0
	rd = render.New() // rd 정의

	mux := pat.New()
	mux.Get("/users", getUsersHandler)
	mux.Post("/users", addUsersHandler)
	mux.Get("/hello", templateHandler)
	mux.Get("/", indexHandler)
	return mux
}
