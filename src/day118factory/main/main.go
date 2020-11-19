package main

import (
	"day118factory/model"
	"fmt"
)

//golang 中的结构体没有构造函数，使用工厂模式来解决这个问题

func main() {
	var stu model.Student
	stu.Name = "tom"

	fmt.Println(stu.Name)

	// 因为student2 首字母是小写，通过工厂模式引用
	var stu2 = model.NewStudent("jack", 66.2)
	fmt.Println(stu2.GetName())
	stu2.SetName("tt")
	fmt.Println(stu2.GetName())

}
