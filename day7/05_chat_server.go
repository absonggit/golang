package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

// 使用02_tcp_client.go 做客户端验证
type Client struct {
	C    chan string // 用户发送数据管道
	Name string      // 用户名
	Addr string      // 网络地址
}

// 保存在线用户
var OnlineMap map[string]Client

var message = make(chan string)

func Manager() {
	// 给map分配空间
	OnlineMap = make(map[string]Client)

	for {
		msg := <-message
		for _, cli := range OnlineMap {
			cli.C <- msg
		}
	}
}

func WriteMsgToClient(cli Client, conn net.Conn) {
	for msg := range cli.C {
		conn.Write([]byte(msg + "\n"))
	}

}
func MakeMsg(cli Client, msg string) (buf string) {
	buf = "[" + cli.Addr + "]" + cli.Name + ": " + msg
	return
}

// 处理用户连接
func HandleConn(conn net.Conn) {
	defer conn.Close()
	// 获取客户端网络地址
	cliAddr := conn.RemoteAddr().String()
	// 创建结构体
	cli := Client{make(chan string), cliAddr, cliAddr}
	// 添加到map
	OnlineMap[cliAddr] = cli
	//  给客户端发送信息
	go WriteMsgToClient(cli, conn)
	// 广播在线
	message <- MakeMsg(cli, "login")
	cli.C <- MakeMsg(cli, "I am here")
	isQuit := make(chan bool)
	hasData := make(chan bool)
	// 接收用户发送的数据
	go func() {
		buf := make([]byte, 2048)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				isQuit <- true
				fmt.Println("Error: ", err)
				return
			}
			msg := string(buf[:n-2])
			if len(msg) == 3 && msg == "who" {
				conn.Write([]byte("User list:\n"))
				for _, tmp := range OnlineMap {
					msg = tmp.Addr + ":" + tmp.Name + "\n"
					conn.Write([]byte(msg))
				}
			} else if len(msg) >= 8 && msg[:6] == "rename" {
				name := strings.Split(msg, " ")[1]
				cli.Name = name
				OnlineMap[cliAddr] = cli
				conn.Write([]byte("Rename ok\n"))
			} else if len(msg) == 4 && msg[:4] == "exit" {
				isQuit <- true
			} else {
				if len(msg) == 1 {
					continue
				}
				message <- MakeMsg(cli, msg)
			}
			hasData <- true
		}
	}()

	for {
		// 通过select检测channel的流动
		select {
		case <-isQuit:
			delete(OnlineMap, cliAddr) // 当前用户从map移除
			message <- MakeMsg(cli, "login out")
		case <-hasData:

		case <-time.After(60 * time.Second):
			delete(OnlineMap, cliAddr) // 当前用户从map移除
			message <- MakeMsg(cli, "time out leave out")
			return
		}
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

	// 转发消息协程，遍历map 给每个成员发送消息
	go Manager()

	// 主协程，循环阻塞等待用户连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		go HandleConn(conn) // 处理用户连接
	}
}
