package main

import (
	"fmt"
	"net"
	"strings"
)

func HandleConn(conn net.Conn) {
	// 函数调用结束，关闭连接
	defer conn.Close()

	// 获取客户端网络地址信息
	addr := conn.RemoteAddr().String()
	fmt.Printf("[%s]: addr conncet sucessful\n", addr)
	for {
		// 读取用户数据
		buf := make([]byte, 2048)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("[%s]: %s\n", addr, string(buf[:n]))
		if "exit" == string(buf[:n-2]) {
			fmt.Printf("[%s]: exit\n", addr)
			return
		}

		// 把数据转换成大写
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}

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

		// 开启协程处理用户请求
		go HandleConn(conn)
	}
}
