package main

import (
	"fmt"
)

func main() {

	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	s := "Hello,World"
	fmt.Println("我在这里")

	for i := 0; i < len(s); i++ {
		fmt.Printf("%d %c\n", i, s[i])
	}

	for i, data := range s {
		fmt.Printf("%d %c\n", i, data)
	}
	// break 用于for、switch、select， continue 仅用于for循环
	// break 跳出当前循环，continue是跳至下次循环
	// goto 可以用在任何地方，但不能跨函数使用。跳至标签位置，不建议使用

	for i := 1; i <= 10; i++ {
		if i >= 5 && i <= 7 {
			continue
			//break
		}
		fmt.Println(i)
	}

	fmt.Println("cccccccccc")
	goto End
	fmt.Println("bbbbbbbbbb")
End:
	fmt.Println("aaaaaaaaaa")
}
