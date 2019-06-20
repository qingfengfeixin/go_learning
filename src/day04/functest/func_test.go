package functest

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValue() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func TestFn(t *testing.T) {
	a, _ := returnMultiValue() // 使用 _ 忽略两个返回值中的一个
	t.Log(a)
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func SlowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn1(t *testing.T) {
	tsSf := timeSpent(SlowFun)
	t.Log(tsSf(10))
}

func Sum(ops ...int) int {
	// 可变长参数
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestFn2(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5))
}

func Clear() {
	fmt.Println("Clear resource.")
}

func TestDefer(t *testing.T) {
	defer Clear() //defer 可以用于程序释放资源
	fmt.Println("Start")
	panic("Fatal error")
	//fmt.Println("End.")  // panic后的代码不可达到
}
