package fib

// 变量的赋值

import "testing"

func TestFibList(t *testing.T) {
	a := 1
	var tmp int
	t.Log(a)
	for i := 0; i < 10; i++ {
		a, tmp = tmp+a, a
		t.Log(a)
	}

}
