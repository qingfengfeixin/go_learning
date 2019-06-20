package operator

import "testing"

const (
	Readable = 1 << iota
	Writeable
	Executable
)

func TestCompareArry(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 2, 5}
	//c := [...]int{1, 2, 3, 4, 5}   //个数不对无法比较会报错
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	//t.Log(a == c)
	t.Log(a == d)
}

func TestBitClear(t *testing.T) {
	t.Log(Readable, Writeable, Executable)
	a := 7 //0111
	t.Log(a&Readable == Readable, a&Writeable == Writeable, a&Executable == Executable)
	a = a &^ Readable //将可读位清除
	t.Log(a, a&Readable == Readable, a&Writeable == Writeable, a&Executable == Executable)
}
