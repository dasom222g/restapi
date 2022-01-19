package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/antage/eventsource"
	"github.com/gorilla/pat"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

var rd *render.Render

var currentID int
var userMap map[int]*User
var sendMessage chan SendMessageInfo

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type SendMessageInfo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		rd.Text(w, http.StatusBadRequest, err.Error())
	}
	// user정보 새팅
	currentID++
	user.ID = currentID
	user.CreatedAt = time.Now()
	userMap[user.ID] = user

	// 입장 메시지
	messageInfo := &SendMessageInfo{
		ID:        user.ID,
		Name:      user.Name,
		Message:   fmt.Sprintf("%s 님이 입장하셨습니다.", user.Name),
		CreatedAt: time.Now(),
	}
	log.Println("messageInfo", messageInfo)
	setSendMessage(messageInfo)
	rd.JSON(w, http.StatusOK, &user)
}

func handleGetUsers(w http.ResponseWriter, _ *http.Request) {
	if len(userMap) == 0 {
		rd.JSON(w, http.StatusOK, "No user")
		return
	}

	users := []*User{}
	for _, value := range userMap {
		users = append(users, value)
	}
	rd.JSON(w, http.StatusOK, users)
}

func handlePostMessage(w http.ResponseWriter, r *http.Request) {
	messageInfo := new(SendMessageInfo)
	err := json.NewDecoder(r.Body).Decode(&messageInfo)
	if err != nil {
		rd.Text(w, http.StatusBadRequest, err.Error())
		return
	}

	messageInfo.CreatedAt = time.Now()
	setSendMessage(messageInfo)
}

func setSendMessage(messageInfo *SendMessageInfo) {
	sendMessage <- *messageInfo
	// sendMessage <- SendMessageInfo{
	// 	messageInfo.ID,
	// 	messageInfo.Name,
	// 	messageInfo.Message,
	// 	messageInfo.CreatedAt,
	// }
}

func processSendMessage(es eventsource.EventSource) {
	log.Println("len!!!", len(sendMessage))
	for messageInfo := range sendMessage {
		data, _ := json.Marshal(&messageInfo)
		es.SendEventMessage(string(data), "", strconv.Itoa(time.Now().Nanosecond()))
	}
}

func InitData() {
	rd = render.New()
	currentID = 0
	userMap = make(map[int]*User)
	sendMessage = make(chan SendMessageInfo)
}

func NewHttpHandler() http.Handler {
	// InitData()

	es := eventsource.New(nil, nil)
	defer es.Close()

	mux := pat.New()
	mux.Handle("/stream", es) // es 오픈될때 매핑
	mux.Post("/user", handleCreateUser)
	mux.Get("/users", handleGetUsers)
	mux.Post("/message", handlePostMessage)

	go processSendMessage(es) // 메시지가 들어오면 모든 클라이언트에게 알림보냄
	n := negroni.Classic()
	n.UseHandler(mux)

	return n
}
