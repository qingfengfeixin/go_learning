package typet

import "testing"

type MyInt int64

func TestImplicit(t *testing.T) {
	var (
		a int = 1
		b int
		c MyInt
	)
	b = a //连int64 也不能进行赋值，必须类型一样，或者通过显式的类型转换，不支持隐式转换
	c = MyInt(b)
	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1 // 指针
	aPtr := &a
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string //默认初始化为空字符串
	t.Log("*" + s + "*")
	t.Log(len(s))

}
