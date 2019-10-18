package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 写文件
func WriteFile(path string) {
	// 打开文件
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	// 关闭文件
	defer f.Close()
	var buf string
	for i := 0; i < 10; i++ {
		buf = fmt.Sprintf("i = %d\n", i)

		_, err := f.WriteString(buf)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
	}
}

// 读文件
func ReadFile(path string) {
	// 打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	// 关闭文件
	defer f.Close()
	buf := make([]byte, 1024*2) // 2k大小
	n, err := f.Read(buf)
	if err != nil && err != io.EOF { // 文件出错，同时没有到结尾
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(string(buf[:n]))
}

// 每次读一行
func ReadFileLine(path string) {
	// 打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	// 关闭文件
	defer f.Close()

	// 创建缓冲区，把内容放入缓冲区
	r := bufio.NewReader(f)
	for {
		// 遇到"\n"结束读取
		buf, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error: ", err)
		}
		fmt.Printf(string(buf))
	}
}

func main() {
	path := "./demo.txt"
	WriteFile(path)
	ReadFile(path)
	ReadFileLine(path)
}
