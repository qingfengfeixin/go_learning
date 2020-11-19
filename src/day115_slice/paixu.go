package main

import (
	"fmt"
	"math/rand"
	"time"
)

func data(n int) []int {

	var s1 []int = make([]int, n)
	rand.Seed(time.Now().UnixNano())
	len := len(s1)

	for i := 0; i < len; i++ {
		s1[i] = rand.Intn(100)
	}
	fmt.Print(s1, "\n")

	return s1

}

func maopao(s []int) {
	//冒泡排序
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s)-i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	fmt.Print(s, "\n")
}

func paixu(n int) {
	s := data(n)
	maopao(s)
	time.Sleep(1 * time.Nanosecond)
	s = data(n)
	maopao(s)

}
