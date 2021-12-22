package check

import (
	"fmt"
	"log"
	"net/http"
)

func CheckError(err error, w http.ResponseWriter, code int) {
	if err != nil {
		w.WriteHeader(code)
		fmt.Fprint(w, err)
		log.Panic(err)
	}
}
