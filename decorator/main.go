package main

import (
	"fmt"

	"github.com/dasom222g/restapi/decorator/check"
	"github.com/tuckersGo/goWeb/web9/cipher"
	"github.com/tuckersGo/goWeb/web9/lzw"
)

type Component interface {
	Operation(string)
}

var sendData string    // 전송 데이터
var recevieData string // 수신 데이터

// 메인 기능 (send)
type SendComponent struct{}

func (s SendComponent) Operation(data string) {
	sendData = data
}

// 부가 기능 (send)
type ZipComponent struct {
	com Component
}

func (self *ZipComponent) Operation(data string) {
	// 압축 기능 구현
	zipData, err := lzw.Write([]byte(data))
	check.CheckError(err)

	self.com.Operation(string(zipData))
}

type EncrytComponent struct {
	key string
	com Component
}

func (self *EncrytComponent) Operation(data string) {
	// 암호화 기능 구현
	encryptData, err := cipher.Encrypt([]byte(data), self.key)
	check.CheckError(err)

	self.com.Operation(string(encryptData))
}

// 메인 기능 (receive)
type ReceiveComponent struct{}

func (self *ReceiveComponent) Operation(data string) {
	recevieData = data
}

// 부가 기능 (receive)
type UnZipComponent struct {
	com Component
}

func (self *UnZipComponent) Operation(data string) {
	unzipData, err := lzw.Read([]byte(data))
	check.CheckError(err)
	self.com.Operation(string(unzipData))
}

type DecryptComponent struct {
	key string
	com Component
}

func (self *DecryptComponent) Operation(data string) {
	decrytData, err := cipher.Decrypt([]byte(data), self.key)
	check.CheckError(err)
	self.com.Operation(string(decrytData))
}

func main() {
	// 순서: 암호화 -> 압축
	sender := &EncrytComponent{
		key: "abcd",
		com: &ZipComponent{
			com: &SendComponent{},
		},
	}
	sender.Operation("Hello world!")
	fmt.Println(sendData)

	// 순서: 압축 풀기 -> 복호화
	receiver := &UnZipComponent{
		com: &DecryptComponent{
			key: "abcd",
			com: &ReceiveComponent{},
		},
	}

	receiver.Operation(sendData)
	fmt.Println(recevieData)
}
