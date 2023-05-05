package main

import (
	"fmt"
	"github.com/xhyonline/lfutil/request"
)

func main() {
	fmt.Println(request.Get("https://www.baidu.com"))
}
