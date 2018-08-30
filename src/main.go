package main

import (
	"fmt"
	"net"
	"os"
	stProto "../proto"

	"github.com/gogo/protobuf/proto"
)

func HandleConn(conn net.Conn) {
	defer conn.Close()

	addr := conn.RemoteAddr().String()
	fmt.Println(addr, " connect sucessful.")

	buf := make([]byte, 4096, 4096)
	for {
		cnt, err := conn.Read(buf)
		if err != nil {
			panic(err)
		}

		stReceive := &stProto.UserInfo{}
		pData := buf[:cnt]

		//protobuf解码
		err = proto.Unmarshal(pData, stReceive)
		if err != nil {
			panic(err)
		}

		fmt.Println("receive ", conn.RemoteAddr(), stReceive)
		if stReceive.Message == "stop" {
			os.Exit(1)
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", "172.18.62.46:6600")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		//处理用户请求，新建一个协程
		go HandleConn(conn)
	}
}
