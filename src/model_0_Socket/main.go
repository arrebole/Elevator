package main

import (
	"fmt"
	"net"
)

func main() {
	// 1、监听端口
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		return
	}

	for {
		// 2、接受请求
		conn, err := listen.Accept()
		if err != nil {
			return
		}
		go process(conn)
	}

}

// 3 处理请求
func process(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		if err != nil {
			return
		}
		fmt.Println(string(buf))
	}
}
