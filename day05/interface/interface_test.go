package interface_test

import "testing"

type Programer interface {
	WriteHelloWorld() string
}

// 接口为非侵入式的，实现不需要依赖接口定义
// 类型
type GoProgrammer struct {
}

//数据
func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello world\")"
}

func TestClient(t *testing.T) {
	var p Programer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}
