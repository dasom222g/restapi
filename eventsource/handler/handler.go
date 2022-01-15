package handler

import (
	"net/http"

	"github.com/gorilla/pat"
	"github.com/urfave/negroni"
)

func NewHttpHandler() http.Handler {
	mux := pat.New()
	n := negroni.Classic()
	n.UseHandler(mux)
	return n
}
