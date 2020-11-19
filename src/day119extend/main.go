package main

import "fmt"

// 继承，多个结构体具有相同属性，可以抽象出一个匿名结构体

type Student struct {
	Name  string
	Age   int
	Score int
}

func (stu *Student) ShowInfo() {
	fmt.Println(stu.Name, stu.Age, stu.Score)
}

type pupil struct {
	Student //嵌套匿名结构体
}

type Graduate struct {
	Student
}

func (stu *pupil) Testing() {
	fmt.Println("pupil", stu.Name, "is testing !")
}

func main() {
	var p pupil
	p.Name = "tom"
	p.ShowInfo()
	p.Testing()

}
