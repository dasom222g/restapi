package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/pat"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
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

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	/*
		data, _ := json.Marshal(users)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(data))
	*/
	// render로 소스 축약
	fmt.Println("invoke!!")
	rd.JSON(w, http.StatusOK, users)
}

func addUserHandler(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		/*
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, err)
		*/
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
	/*
		tmpl, err := template.New("tmpl").ParseFiles("templates/hello.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err)
			return
		}
		tmpl.ExecuteTemplate(w, "hello.html", "dasomi")
	*/
	rd.HTML(w, http.StatusOK, "info", users)
}

func NewHttpHandler() http.Handler {
	currentID = 0
	// users = make([]*User, 0)
	users = []*User{}
	rd = render.New(render.Options{
		Directory:  "template_contents",
		Extensions: []string{".html", ".tmpl"},
		Layout:     "hello",
	}) // rd 정의

	mux := pat.New()
	n := negroni.Classic() // 새로운 Negroni 생성
	// 미들웨어 스택에 http.Handler 추가. 처리기는 Negroni에 추가된 순서로 호출 (hadler -> logger)
	n.UseHandler(mux)

	mux.Get("/users", getUsersHandler)
	mux.Post("/users", addUserHandler)
	mux.Get("/hello", templateHandler)
	// Negroni생성시 파일서버를 자동으로 지원하므로 public/index.html파일이 있으면 				handler작성 하지않아도 자동으로 해당 파일을 인덱스 페이지로 구동시킴
	// mux.Handle("/", http.FileServer(http.Dir("public")))
	return n
}
