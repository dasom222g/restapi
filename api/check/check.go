package check

import (
	"fmt"
	"net/http"
)

func CheckError(err error, w http.ResponseWriter, code int) bool {
	if err != nil {
		w.WriteHeader(code)
		fmt.Fprint(w, err)
		return true
	}
	return false
}
