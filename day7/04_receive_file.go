package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func ReceiveFile(fileName string, conn net.Conn) {
	// 新建文件
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	buf := make([]byte, 1024*4)

	// 接收文件
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件接收完毕")
			} else {
				fmt.Println("Error:", err)
			}
			return
		}
		f.Write(buf[:n])
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
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	buf := make([]byte, 1024)
	var n int
	n, err = conn.Read(buf) // 读取发送文件名
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fileName := string(buf[:n])
	// 回复"ok"
	conn.Write([]byte("ok"))

	// 接收文件内容
	ReceiveFile(fileName, conn)
}
