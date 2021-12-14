package main

import (
	"net/http"

	"github.com/dasom222g/restapi/fileupload/upload"
)

func main() {
	// http.Handle("/", http.FileServer(http.Dir("public")))
	// http.HandleFunc("/upload", upload.FileUpload)
	http.ListenAndServe(":3000", upload.NewHttpHandler())
}
