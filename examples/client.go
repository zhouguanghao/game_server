package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	stProto "../proto"
	"time"

	"github.com/gogo/protobuf/proto"
)

func main() {
	strIP := "172.18.62.46:6600"
	var conn net.Conn
	var err error

	//连接服务器
	for conn, err = net.Dial("tcp", strIP); err != nil; conn, err = net.Dial("tcp", strIP) {
		fmt.Println("connect ", strIP, " fail")
		time.Sleep(time.Second)
		fmt.Println("reconnect...")
	}
	fmt.Println("connect ", strIP, " sucess")
	defer conn.Close()

	//发送消息
	cnt := 0
	sender := bufio.NewScanner(os.Stdin)
	for sender.Scan() {
		cnt++
		stSend := &stProto.UserInfo{
			Message: sender.Text(),
			Length:  *proto.Int(len(sender.Text())),
			Cnt:     *proto.Int(cnt),
		}

		//protobuf编码
		pData, err := proto.Marshal(stSend)
		if err != nil {
			panic(err)
		}

		//发送
		conn.Write(pData)
		if sender.Text() == "stop" {
			return
		}
	}
}
