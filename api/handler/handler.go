package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dasom222g/restapi/api/check"
	"github.com/gorilla/mux"
)

var currentId int
var userMap map[int]*user // 생성된 user 정보 저장(Go에 저장되는 것으로 run 끄면 reset)

type user struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreateAt  time.Time `json:"create_at"`
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	if len(userMap) == 0 {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No user")
		return
	}
	fmt.Println(userMap)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	for _, val := range userMap {
		data, _ := json.Marshal(val)
		fmt.Fprint(w, string(data), ",")
	}
}

func handleGetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	check.CheckError(err, w, http.StatusBadRequest)

	// map에서 유효하지 않은 값을 요청하면 두번째 리턴값으로 boolean을 보냄
	user, exists := userMap[id]
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Not found Id: ", id)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))
}

func handleCreatUser(w http.ResponseWriter, r *http.Request) {
	// user 생성
	user := new(user)
	err := json.NewDecoder(r.Body).Decode(user)
	check.CheckError(err, w, http.StatusBadRequest)

	// create user setting
	currentId++
	id := currentId
	user.ID = id
	user.CreateAt = time.Now()
	userMap[id] = user

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))

}

func handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	check.CheckError(err, w, http.StatusBadRequest)

	// 요청한 id의 데이터가 있는지 확인하여 없으면 No user , 있으면 Deleted
	_, exists := userMap[id]
	if !exists {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "No user Id : ", id)
		return
	}

	delete(userMap, id)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted success ", id)
}

func NewHttpHandler() http.Handler {
	// 초기화
	currentId = 0
	userMap = make(map[int]*user)

	mux := mux.NewRouter()
	// mux := http.NewServeMux()
	mux.HandleFunc("/", handleIndex)
	mux.HandleFunc("/users", handleUsers).Methods("GET")
	mux.HandleFunc("/users", handleCreatUser).Methods("POST")
	mux.HandleFunc("/users/{id:[0-9]+}", handleGetUser).Methods("GET")
	mux.HandleFunc("/users/{id:[0-9]+}", handleDeleteUser).Methods("DELETE")
	return mux
}
