package main

import (
	"fmt"
	"net/http"
	"time"
)

// w 给客户端回复数据 req 读取客户端发送的数据
func HandConn(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hello go"))
	fmt.Println(time.Now().UTC(), req.RemoteAddr, req.Method, req.URL, req.UserAgent())
}
func main() {
	http.HandleFunc("/", HandConn)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

}
