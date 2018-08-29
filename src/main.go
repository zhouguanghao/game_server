package main

import (
	"fmt"
	"net"
	"strings"
)

func HandleConn(conn net.Conn) {
	defer conn.Close()

	addr := conn.RemoteAddr().String()
	fmt.Println(addr, " connect sucessful.")

	buf := make([]byte, 2049)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("err=", err)
			return
		}
		fmt.Printf("[%s]: %s\n", addr, string(buf[:n]))
		fmt.Println("len = ", len(string(buf[:n])))

		if "exit" == string(buf[:n-1]) {
			fmt.Println(addr, " exit")
			return
		}

		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8005")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("err=", err)
			return
		}
		//处理用户请求，新建一个协程
		go HandleConn(conn)
	}
}
