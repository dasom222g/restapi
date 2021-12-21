package check

import (
	"fmt"
	"net/http"
)

func CheckError(err error, w http.ResponseWriter, code int) {
	if err != nil {
		w.WriteHeader(code)
		fmt.Fprint(w, err)
	}
}
