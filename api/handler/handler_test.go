package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var contentType string = "application/json"

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

func TestHandleUsers(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(NewHttpHandler())
	defer ts.Close()

	res, err := http.Get(ts.URL + "/users")
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	data, _ := ioutil.ReadAll(res.Body)
	assert.Contains(string(data), "No user")
}

func TestHandleGetUserNotFound(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(NewHttpHandler())
	defer ts.Close()

	res, err := http.Get(ts.URL + "/users/58")
	assert.NoError(err)
	assert.Equal(http.StatusNotFound, res.StatusCode)

	data, _ := ioutil.ReadAll(res.Body)
	assert.Contains(string(data), "Not found", "58")
}

func TestHandleCreatUser(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(NewHttpHandler())
	defer ts.Close()
	s := `{
		"first_name": "kelly",
		"last_name": "dasom",
		"email": "dasom228@gmail.com"
	}`

	// post 요청
	res, err := http.Post(ts.URL+"/users", contentType, strings.NewReader(s))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, res.StatusCode)

	// post 후 response data와 gerUser 의 response 데이터 비교
	createUser := new(user)
	err = json.NewDecoder(res.Body).Decode(createUser)
	assert.NoError(err)

	res, err = http.Get(ts.URL + "/users/" + strconv.Itoa(createUser.ID))
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	getUser := new(user)
	err = json.NewDecoder(res.Body).Decode(getUser)
	assert.NoError(err)

	// 비교
	assert.Equal(createUser.ID, getUser.ID)
	assert.Equal(createUser.FirstName, getUser.FirstName)
}

func TestDeleteUserNotFound(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(NewHttpHandler())
	defer ts.Close()

	req, _ := http.NewRequest("DELETE", ts.URL+"/users/1", nil)
	res, err := http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusBadRequest, res.StatusCode)

	data, _ := ioutil.ReadAll(res.Body)
	assert.Contains(string(data), "No user Id : 1")
}

func TestDeleteUser(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(NewHttpHandler())
	defer ts.Close()

	// post
	s := `{
		"first_name": "Test01",
		"last_name": "Test_01",
		"email": "Test01@gmail.com"
	}`
	res, err := http.Post(ts.URL+"/users", contentType, strings.NewReader(s))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, res.StatusCode)

	user := new(user)
	err = json.NewDecoder(res.Body).Decode(user)
	assert.NoError(err)
	id := user.ID

	// delete testing
	req, _ := http.NewRequest("DELETE", ts.URL+"/users/"+strconv.Itoa(id), nil)
	res, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	data, _ := ioutil.ReadAll(res.Body)
	log.Println(string(data))
	assert.Equal(string(data), "Deleted success "+strconv.Itoa(id))
}
