package main

// 闭包的演示

import (
	"fmt"
	"strings"
)

func addupper() func(int) int {
	var n int = 10 //n初始化一次
	return func(x int) int {
		n += x
		return n
	}
}

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func makeSuffix1(suffix string, name string) string {
	if !strings.HasSuffix(name, suffix) {
		return name + suffix
	}
	return name
}

func main() {
	f := addupper()
	fmt.Print(f(1), "\n") // 11
	fmt.Print(f(2), "\n") // 13
	fmt.Print(f(3), "\n") // 16

	f1 := addupper()
	fmt.Print(f1(1), "\n") // 11
	fmt.Print(f1(2), "\n") // 13

	f2 := makeSuffix(".jpg")
	fmt.Println("文件名处理后", f2("gorun"))
	fmt.Println("文件名处理后", f2("go2.jpg"))

	fmt.Println("文件名处理后", makeSuffix1(".jpg", "go3"))

}
