package handler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleIndex(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(NewHttpHandler()) // 테스트 서버(mock-up서버 생성)
	defer ts.Close()

	// ts.URL : http://127.0.0.1:60624 (index url)
	res, err := http.Get(ts.URL)
	// fmt.Println(ts.URL)
	assert.NoError(err)

	assert.Equal(http.StatusOK, res.StatusCode)

	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello, world!", string(data))

}
