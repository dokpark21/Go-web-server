package main

import (
	"fmt"

	"github.com/tuckersGo/goWeb/web9/cipher"
	"github.com/tuckersGo/goWeb/web9/lzw"
)

// Decorator Pattern
// Data -> Encrypt -> Zip -> Send
// Receive -> Unzip -> Decrypt -> Data

type Component interface {
	Operator(string)
}

var sendData string
var recvData string

type SendComponent struct {
}

func (self *SendComponent) Operator(data string) {
	// Send Data
	sendData = data
}

type ZipComponent struct {
	com Component
}

func (self *ZipComponent) Operator(data string) {
	zipData, err := lzw.Write([]byte(data))
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(zipData))
}

type EncryptComponent struct {
	key string
	com Component
}

func (self *EncryptComponent) Operator(data string) {
	encryptData, err := cipher.Encrypt([]byte(data), self.key)
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(encryptData))
}

type DecryptComponent struct {
	key string
	com Component
}

func (self *DecryptComponent) Operator(data string) {
	decryptData, err := cipher.Decrypt([]byte(data), self.key)
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(decryptData))
}

type UnZipComponent struct {
	com Component
}

func (self *UnZipComponent) Operator(data string) {
	unzipData, err := lzw.Read([]byte(data))
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(unzipData))
}

type ReadComponent struct {
}

func (self *ReadComponent) Operator(data string) {
	recvData = data
}

func main() {
	// EncryptComponent가 ZipComponent를 가져오고, 그것이 SendComponent를 가져온다.
	sender := &EncryptComponent{key: "abcdef", com: &ZipComponent{com: &SendComponent{}}}
	sender.Operator("Hello World")
	fmt.Println(sendData)
	// 압축해제 -> 복호화 -> 데이터
	receiver := &UnZipComponent{com: &DecryptComponent{key: "abcdef", com: &ReadComponent{}}}
	receiver.Operator(sendData)
	fmt.Println(recvData)
}
