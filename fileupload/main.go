package main

import (
	"net/http"

	"github.com/dasom222g/restapi/fileupload/upload"
)

func main() {
	http.ListenAndServe(":3000", upload.NewHttpHandler())
}
