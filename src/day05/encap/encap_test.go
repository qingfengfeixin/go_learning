package encap

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"0", "Jack", 18}
	e1 := Employee{Name: "Mike", Age: 20}
	e1.Id = "1"
	e2 := new(Employee) // 返回指针
	e2.Id = "2"
	e2.Name = "Tom"
	t.Log(e, e1, e2)
	t.Logf("e is %T", e)   //e和e1是实例
	t.Logf("e2 is %T", e2) //e2是一个指针引用,相当于以下的方式
	t.Logf("e is %T", &e)
}

//对比以下两种封装方式，推荐 方式1
//方式1
// func (e *Employee) String() string {
// 	fmt.Printf("Address is %x", unsafe.Pointer(&e.Name))
// 	return fmt.Sprintf("ID:%s/Name:%s/Age:%d", e.Id, e.Name, e.Age)
// }

// 方式2
func (e Employee) String() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func TestStructOperations(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	t.Log(e.String())
}
