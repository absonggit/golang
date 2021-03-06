package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// 连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("服务器连接成功")
	msg := `    who 获取当前用户列表
    rename <newname> 修改当前用户名
    exit 退出`
	fmt.Println(msg)
	defer conn.Close()

	// 输入内容，发送至服务器
	go func() {
		str := make([]byte, 2048)
		for {
			n, err := os.Stdin.Read(str)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			conn.Write(str[:n])
		}
	}()

	// 接收服务器回复数据
	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
