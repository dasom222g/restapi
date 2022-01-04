package main

import (
	"net/http"

	"github.com/dasom222g/restapi/decoratorHttpHandle/httpHandler"
)

func main() {
	http.ListenAndServe(":3000", httpHandler.NewHttpHandler())
}
