package main

import "fmt"

func main() {
	//map的声明不分配空间，初始化需要make，分配内存后才能赋值和使用
	var a map[int]string
	a = make(map[int]string)
	a[0] = "a"
	a[1] = "b"
	a[2] = "c"
	for i, v := range a {
		fmt.Println(i, v)
	}
	b := map[int]string{
		1: "a",
		2: "b",
	}
	fmt.Println(b)

	c := make(map[int]string)
	c[1] = "a"
	delete(a, 2)

	//map中value是map需要make初始化
	student := make(map[string]map[string]string)
	student["stu01"] = make(map[string]string)
	student["stu01"]["name"] = "tom"
	student["stu01"]["sex"] = "male"

	student["stu02"] = make(map[string]string)
	student["stu02"]["name"] = "lucy"
	student["stu02"]["sex"] = "female"

	for i, v := range student {
		fmt.Println(i, v)
	}

	//monster()
	//paixu()
	stude()

}
