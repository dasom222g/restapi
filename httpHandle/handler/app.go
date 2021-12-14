package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

type User struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Age       int       `json:"age"`
	CreateAt  time.Time `json:"create_at"`
}

// struct
type UserInfo struct{}

func (u *UserInfo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)                           // user 생성
	err := json.NewDecoder(r.Body).Decode(user) // json형식으로 요청받은 데이터를 go형식으로 변환
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "BadRequest: ", err)
		return
	}
	user.CreateAt = time.Now()
	data, _ := json.Marshal(user) // json형식으로 변환(byte[], error 리턴)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data)) // response 보냄
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s!", name)
}

func NewHttpHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", IndexHandler)

	mux.HandleFunc("/bar", barHandler)

	mux.Handle("/userinfo", &UserInfo{})
	return mux
}
