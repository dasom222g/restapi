package handler

import (
	"encoding/json"
	"fmt"
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

func addUser(url string, assert *assert.Assertions, appendStr string) *user {
	s := fmt.Sprintf(`{
		"first_name": "kelly%s",
		"last_name": "dasom%s",
		"email": "dasom228@gmail.com%s"
	}`, appendStr, appendStr, appendStr)

	// log.Println(s)

	// post 요청
	res, err := http.Post(url, contentType, strings.NewReader(s))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, res.StatusCode)

	// response 데이터 반환
	createdUser := new(user)
	err = json.NewDecoder(res.Body).Decode(createdUser)
	assert.NoError(err)

	return createdUser
}

/* Test code*/
func TestHandleIndex(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(NewHttpHandler()) // 테스트 서버(mock-up서버 생성)
	defer ts.Close()

	res, err := http.Get(ts.URL)
	assert.NoError(err)

	assert.Equal(http.StatusOK, res.StatusCode)

	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello, world!", string(data))

}

func TestHandleGetUsersNotFound(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(NewHttpHandler())
	defer ts.Close()

	res, err := http.Get(ts.URL + "/users")
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	data, _ := ioutil.ReadAll(res.Body)
	assert.Contains(string(data), "No user")
}

func TestHandleGetUsers(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(NewHttpHandler())
	defer ts.Close()

	createdUser1 := addUser(ts.URL+"/users", assert, "11")
	createdUser2 := addUser(ts.URL+"/users", assert, "22")

	res, err := http.Get(ts.URL + "/users")
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	users := []*user{}
	err = json.NewDecoder(res.Body).Decode(&users)
	assert.NoError(err)
	assert.NotZero(len(users))
	assert.Equal(2, len(users))

	assert.Equal(createdUser1.ID, users[0].ID)
	assert.Equal(createdUser2.ID, users[1].ID)
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

	// post 후 response 데이터와 gerUser 의 response 데이터 비교
	createdUser := addUser(ts.URL+"/users", assert, "11")

	res, err := http.Get(ts.URL + "/users/" + strconv.Itoa(createdUser.ID))
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	getUser := new(user)
	err = json.NewDecoder(res.Body).Decode(getUser)
	assert.NoError(err)

	// 비교
	assert.Equal(createdUser.ID, getUser.ID)
	assert.Equal(createdUser.FirstName, getUser.FirstName)
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
	createdUser := addUser(ts.URL+"/users", assert, "11")
	id := createdUser.ID

	// delete testing
	req, _ := http.NewRequest("DELETE", ts.URL+"/users/"+strconv.Itoa(id), nil)
	res, err := http.DefaultClient.Do(req)
	assert.NoError(err)
	data, _ := ioutil.ReadAll(res.Body)
	log.Println(string(data))
	assert.Equal(string(data), "Deleted success "+strconv.Itoa(id))
}

func TestUpdateUserNotFound(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(NewHttpHandler())
	defer ts.Close()

	s := `{
		"first_name": "update kelly",
		"last_name": "update dasom",
		"email": "dasom228@gmail.com"
	}`

	req, _ := http.NewRequest("PUT", ts.URL+"/users/1", strings.NewReader(s))
	res, err := http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusBadRequest, res.StatusCode)

	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("No user Id : 1", string(data))
}

func TestUpdateUser(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(NewHttpHandler())
	defer ts.Close()

	createdUser := addUser(ts.URL+"/users", assert, "")
	assert.NotEqual(0, createdUser.ID)

	s := `{
		"first_name": "update kelly",
		"last_name": "update dasom"
	}`

	req, _ := http.NewRequest("PUT", ts.URL+"/users/"+strconv.Itoa(createdUser.ID), strings.NewReader(s))
	res, err := http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	updatedUser := new(user)
	err = json.NewDecoder(res.Body).Decode(updatedUser)
	assert.NoError(err)
	assert.Equal(createdUser.ID, updatedUser.ID)
	assert.NotEqual(createdUser.FirstName, updatedUser.FirstName)
	assert.NotEqual(createdUser.LastName, updatedUser.LastName)
	assert.Equal(createdUser.Email, updatedUser.Email)
}
