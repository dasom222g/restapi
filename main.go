package main

import (
	"net/http"

	"github.com/dasom222g/restapi/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}
