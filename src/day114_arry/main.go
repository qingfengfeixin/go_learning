package main

//数组相关

import (
	"fmt"
	"math/rand"
	"time"
)

func test01(arr [3]int) {
	arr[0] = 100
	arr[1] = 200
	arr[2] = 300
}

func test02(arr *[3]int) {
	(*arr)[0] = 100
	(*arr)[1] = 200
	(*arr)[2] = 300
}

func main() {
	//数组是连续存放的
	var a [3]int
	fmt.Printf("a数组的地址 %p \n第一个元素地址 %p \n第二个元素地址 %p \n", &a, &a[0], &a[1])

	var b = [...]int{1, 2, 3}

	test01(b)
	for i, v := range b {
		fmt.Printf("b[%v] = %v \n", i, v)
	}

	// 函数中修改数组入参

	test02(&b)
	for i, v := range b {
		fmt.Printf("b[%v] = %v \n", i, v)
	}

	//倒序数组元素

	var (
		c   [5]int
		len int = len(c)
	)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len; i++ {
		c[i] = rand.Intn(100)
	}
	fmt.Print(c)

	for i := 0; i < len/2; i++ {
		c[i], c[len-i-1] = c[len-i-1], c[i]
	}
	fmt.Print(c)

}
