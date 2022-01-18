package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/antage/eventsource"
	"github.com/gorilla/pat"
	"github.com/urfave/negroni"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func addUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	fmt.Fprint(w, user.Name)
}

func NewHttpHandler() http.Handler {
	es := eventsource.New(nil, nil)
	defer es.Close()

	mux := pat.New()
	mux.Handle("/stream", es) // es 오픈될때 매핑
	mux.Post("/user", addUser)

	n := negroni.Classic()
	n.UseHandler(mux)

	return n
}
