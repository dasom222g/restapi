package main

import (
	"net/http"

	"github.com/dasom222g/restapi/api/handler"
)

func main() {
	http.ListenAndServe(":3000", handler.NewHttpHandler())
}
