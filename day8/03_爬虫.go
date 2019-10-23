package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()
	buf := make([]byte, 1024*4)
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		result += string(buf[:n])
	}
	return
}

func SpiderPape(i int, page chan int) {
	url := "https://tieba.baidu.com/f?kw=%E8%8B%B1%E9%9B%84%E8%81%94%E7%9B%9F&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
	fmt.Printf("正在爬第%02d个页面\n", i)
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fileName := strconv.Itoa(i) + ".html"
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	f.WriteString(result)
	f.Close()
	page <- i
}
func DoWork(start, end int) {
	fmt.Printf("从%02d 到 %02d 开始爬取数据\n", start, end)
	page := make(chan int)
	for i := start; i <= end; i++ {
		go SpiderPape(i, page)
	}
	for i := start; i <= end; i++ {
		fmt.Printf("第%02d个页面爬取完成\n", <-page)
	}
}
func main() {
	var start, end int
	fmt.Printf("输入起始页：")
	fmt.Scan(&start)
	fmt.Println()
	fmt.Printf("输入结束页：")
	fmt.Scan(&end)
	fmt.Println()
	DoWork(start, end)
}
