package main

import "fmt"

type stu struct {
	Name string
	Age  int
	Addr string
}

func stude() {
	students := make(map[string]stu, 10)
	stu1 := stu{"tom", 18, "北京"}
	stu2 := stu{"mary", 14, "上海"}

	students["No1"] = stu1
	students["No2"] = stu2

	fmt.Println(students)

	for k, v := range students {
		fmt.Println(k, v, v.Name, v.Age, v.Addr)
	}

}
