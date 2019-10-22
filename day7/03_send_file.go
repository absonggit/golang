// 03_send_file.go
package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func SendFile(path string, conn net.Conn) {
	// 打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer f.Close()

	buf := make([]byte, 1024*4)
	// 读取文件并发送
	for {
		n, err := f.Read(buf) // 读取内容
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件发送完毕")
			} else {
				fmt.Println("Error: ", err)
			}
			return
		}
		// 发送内容
		conn.Write(buf[:n])
	}
}
func main() {
	// 获取文件信息
	fmt.Printf("输入文件绝对路径：")
	var path string
	fmt.Scan(&path)
	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// 连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer conn.Close()

	// 发送文件名
	_, err = conn.Write([]byte(info.Name()))
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// 接收回复“ok”，可以发送文件
	var n int
	buf := make([]byte, 1024)
	n, err = conn.Read(buf)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	if "ok" == string(buf[:n]) {
		SendFile(path, conn)
	}

}
