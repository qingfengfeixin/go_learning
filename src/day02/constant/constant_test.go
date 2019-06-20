package constant

//常量

import "testing"

// 常量的连续赋值
const (
	Monday = iota + 1
	Tuesday
	Wendnesday
)

const (
	Readable = 1 << iota
	Writeable
	Executable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday, Wendnesday)
}

func TestConstantTry1(t *testing.T) {
	t.Log(Readable, Writeable, Executable)
	a := 5 //0101
	t.Log(a&Readable == Readable, a&Writeable == Writeable, a&Executable == Executable)
}
