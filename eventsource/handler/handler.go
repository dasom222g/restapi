package handler

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/antage/eventsource"
	"github.com/unrolled/render"
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

// func NewHttpHandler() http.Handler {
// 	// InitData()

// 	es := eventsource.New(nil, nil)
// 	defer es.Close()

// 	mux := pat.New()
// 	mux.Handle("/stream", es) // es 오픈될때 매핑

// 	go processSendMessage(es) // 메시지가 들어오면 모든 클라이언트에게 알림보냄
// 	n := negroni.Classic()
// 	n.UseHandler(mux)

// 	return n
// }
