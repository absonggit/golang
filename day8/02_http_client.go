package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://127.0.0.1:8888")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
	buf := make([]byte, 4*1024)
	var tmp string
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("Error:", err)
			break
		}
		tmp += string(buf[:n])
	}
	fmt.Println(tmp)
}
