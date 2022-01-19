package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/antage/eventsource"
	"github.com/dasom222g/restapi/eventsource/check"
	"github.com/gorilla/pat"
	"github.com/urfave/negroni"
)

var currentID int
var userMap map[int]*User

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(&user)
	if check.CheckError(err, w, http.StatusBadRequest) {
		return
	}

	log.Println("id11", currentID)
	currentID++
	log.Println("id22", currentID)
	user.ID = currentID
	log.Println("id33", currentID)
	user.CreatedAt = time.Now()
	userMap[user.ID] = user

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	data, _ := json.Marshal(&user)
	fmt.Fprint(w, string(data))
}

func handleGetUsers(w http.ResponseWriter, _ *http.Request) {
	users := []*User{}
	for _, value := range userMap {
		users = append(users, value)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	data, _ := json.Marshal(&users)
	fmt.Fprint(w, string(data))
}

func NewHttpHandler() http.Handler {
	currentID = 0
	userMap = make(map[int]*User)
	es := eventsource.New(nil, nil)
	defer es.Close()
	log.Println("id!!!!", currentID)

	mux := pat.New()
	mux.Handle("/stream", es) // es 오픈될때 매핑
	mux.Post("/user", handleCreateUser)
	mux.Get("/users", handleGetUsers)

	n := negroni.Classic()
	n.UseHandler(mux)

	return n
}
