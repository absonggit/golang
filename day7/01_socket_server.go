package main

import (
	"fmt"
	"net"
)

func main() {
	// 监听
	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()
	fmt.Println("服务器启动")
	// 阻塞等待用户连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		// 接收用户请求
		buf := make([]byte, 1024)
		n, err1 := conn.Read(buf)
		if err1 != nil {
			fmt.Println("Error:", err1)
		}
		fmt.Println(string(buf[:n]))
		defer conn.Close()
	}
}