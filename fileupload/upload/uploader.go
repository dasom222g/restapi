package upload

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func checkError(err error, w http.ResponseWriter, statusCode int) {
	if err != nil {
		w.WriteHeader(statusCode)
		return
	}
}

func FileUpload(w http.ResponseWriter, r *http.Request) {
	dirName := "./files"
	uploadFile, header, err := r.FormFile("file_upload")
	checkError(err, w, http.StatusBadRequest)
	filePath := fmt.Sprintf("%s/%s", dirName, header.Filename)

	os.MkdirAll(dirName, 0777)       // 폴더 생성
	file, err := os.Create(filePath) // 빈 파일 생성
	defer file.Close()
	checkError(err, w, http.StatusInternalServerError)

	io.Copy(file, uploadFile) // 빈 파일에 업로드된 파일 복사
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, filePath)
}

func NewHttpHandler() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("public")))
	mux.HandleFunc("/upload", FileUpload)
	return mux
}
