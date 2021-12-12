package myapp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathIndex(t *testing.T) {
	assert := assert.New(t) // assert 생성
	// 네트워크를 사용하지 않고 테스트용 response, request 받아오기
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	// IndexHandler 테스트 하기
	IndexHandler(res, req)
	// if res.Code != http.StatusOK {
	// 	t.Fatal("Failed!!", res.Code)
	// }
	// 인자 값을 비교하여 동일하지 않으면 자동으로 fatal시켜줌
	assert.Equal(http.StatusOK, res.Code)

	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello World!", string(data))
}

func TestBarHandler_WithoutName(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)

	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello World!", string(data))
}

func TestBarHandler_WitName(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar?name=kelly", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)

	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello kelly!", string(data))
}

func TestUserInfo_WithoutJson(t *testing.T) {
	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/userinfo", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusBadRequest, res.Code)
}

func TestUserInfo_WithJson(t *testing.T) {
	assert := assert.New(t)
	s := `{
		"first_name": "Kim",
		"last_name": "dasom",
		"age": 32
	}`

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/userinfo", strings.NewReader(s))

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)

	user := new(User)
	// decoder 생성하여 user에 decode한 값 넣음
	err := json.NewDecoder(res.Body).Decode(user)
	assert.Nil(err) // nil이 여야 통과되며 boolean타입 리턴
	assert.Equal("Kim", user.FirstName)
	assert.Equal("dasom", user.LastName)
	assert.Equal(32, user.Age)
}
