package main

import "fmt"

type cat struct {
	Name  string
	Age   int
	Color string
	hobby string
	ptr   *int              // 指针 需要new
	slice []int             // 切片 需要make
	map1  map[string]string //map 需要make
}

func main() {
	var cat1 cat
	fmt.Println(cat1)
	cat1.Name = "小白"
	cat1.Age = 2
	cat1.Color = "white"
	cat1.hobby = "eat"
	fmt.Println(cat1)

	//指针、slice、map的零值都是nil，使用之前需要make分配空间才行
	cat1.ptr = &cat1.Age
	cat1.slice = make([]int, 10)
	cat1.slice[0] = 100
	cat1.map1 = make(map[string]string)
	cat1.map1["key1"] = "tom"
	fmt.Println(cat1)

	//结构体是值类型，传入的时候是值拷贝，互不影响
	cat2 := cat1
	cat2.Age = 200
	fmt.Println("cat1 = ", cat1)
	fmt.Println("cat2 = ", cat2)

	var cat3 cat = cat{}
	cat3.Name = "ly"
	fmt.Println(cat3)

	var cat4 *cat = new(cat) //new出来的是一个指针
	(*cat4).Age = 100
	//也可以这样写,为了使用在底层会对这种写法进行处理，给cat4加上取值（*）运算
	cat4.Age = 200
	fmt.Println(cat4)
	var cat5 *cat = &cat{}
	//也有两种赋值写法
	(*cat5).Age = 300
	cat5.Age = 400
	fmt.Println(cat5)
	point()
	monster()

}
