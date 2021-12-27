package main

import (
	"fmt"

	"github.com/tuckersGo/goWeb/web9/cipher"
	"github.com/tuckersGo/goWeb/web9/lzw"
)

type Component interface {
	Operation(string)
}

var sendData string
var recevieData string

// 메인 기능
type SendComponent struct{}

func (s SendComponent) Operation(data string) {
	sendData = data
}

// 부가 기능
type ZipComponent struct {
	com Component
}

func (self *ZipComponent) Operation(data string) {
	// 압축 기능 구현
	zipData, err := lzw.Write([]byte(data))
	if err != nil {
		panic(err)
	}

	self.com.Operation(string(zipData))
}

type EncrytComponent struct {
	key string
	com Component
}

func (self *EncrytComponent) Operation(data string) {
	// 암호화 기능 구현
	encryptData, err := cipher.Encrypt([]byte(data), self.key)
	if err != nil {
		panic(err)
	}

	self.com.Operation(string(encryptData))
}

func main() {
	sender := EncrytComponent{
		key: "abcd",
		com: &ZipComponent{
			com: &SendComponent{},
		},
	}
	sender.Operation("Hello world!")
	fmt.Println(sendData)
}
