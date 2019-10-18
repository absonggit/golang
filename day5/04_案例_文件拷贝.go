package main

import (
	"fmt"
	"io"
	_ "io"
	"os"
	"path/filepath"
)

func main() {
	list := os.Args
	_, file := filepath.Split(list[0])

	if len(list) != 3 {
		fmt.Printf("Usage:\n     %s <src_file> <dest_file>\n", file)
		return
	}

	if list[1] == list[2] {
		fmt.Printf("Error:\n     The source file and destination file names cannot be the same !\n")
		return
	}

	// 只读方式打开源文件
	Sf, err := os.Open(list[1])
	if err != nil {
		fmt.Printf("Error:\n     %s", err)
		return
	}
	Df, err := os.Create(list[2])
	if err != nil {
		fmt.Printf("Error:\n     %s", err)
		return
	}
	// 关闭文件
	defer Sf.Close()
	defer Df.Close()

	// 从源文件读取内容，写入目的文件，读多少写多少
	buf := make([]byte, 4*1024) // 4k大小
	for {
		n, err := Sf.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("Error:\n     %s", err)
		}
		Df.Write(buf[:n])
	}
}
