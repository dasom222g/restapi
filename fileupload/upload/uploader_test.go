package upload

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileUpload(t *testing.T) {
	assert := assert.New(t)
	uploadFilePath := "C:/Users/USER/dasom/김다솜/dasom.jpg"
	uploadFile, err := os.Open(uploadFilePath)
	assert.NoError(err)
	defer uploadFile.Close()
	buf := &bytes.Buffer{} // NewWriter 에 io.writer로 넣어주기 위한 변수이며 response.body 에 넘겨질 Reader
	// 웹으로 파일을 전송할 때 MIME 포맷을 사용하는데 이를 위한 writer생성
	writer := multipart.NewWriter(buf)
	// CreateFormFile을 이용하여 file생성
	// filepath.Base() : 파일 이름만 잘라줌
	multi, err := writer.CreateFormFile("file_upload", filepath.Base(uploadFilePath))
	assert.NoError(err)
	io.Copy(multi, uploadFile)
	writer.Close()

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/upload", buf)
	// request하는 데이터의 타입을 알려주어 server가 읽을 수 있게함 (필수!!)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	FileUpload(res, req)
	assert.Equal(http.StatusOK, res.Code)

	// 업로한 한 파일과 다운로드 된 파일 비교
	downloadFilePath := "./files/" + filepath.Base(uploadFilePath)
	downloadFile, err := os.Open(downloadFilePath)
	defer downloadFile.Close()
	assert.NoError(err)

	uploadData := []byte{}
	downloadData := []byte{}

	// uploadFile, downloadFile 에 각각 읽어들임
	uploadFile.Read(uploadData)
	downloadFile.Read(downloadData)

	assert.Equal(uploadData, downloadData)

}
