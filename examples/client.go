package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8005")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer conn.Close()

	go func() {
		str := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(str)
			if err != nil {
				fmt.Println("os.Stdin err= ", err)
				return
			}

			//发送到服务器
			conn.Write(str[:n])
		}
	}()

	//接收服务器回复数据
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err=", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
