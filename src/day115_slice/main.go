package main

import "fmt"

// slice相关

func test(s []int) { //引用类型传递的时候 会改变函数实参数
	s[0] = 100
}

func main() {
	// slice and arry
	var arr = [...]int{1, 2, 3, 4, 5, 6}
	s1 := arr[1:3]

	fmt.Printf("\n %p,%p \n", &arr[1], &s1[0])
	fmt.Print(s1, len(s1), cap(s1), "\n")

	fmt.Print(s1, arr, "\n")
	s1[1] = 33
	fmt.Print(s1, arr, "\n")

	s1 = append(s1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12)
	fmt.Print(s1, len(s1), cap(s1))
	fmt.Printf("\n %p,%p \n", &arr[1], &s1[0])

	var s2 []int = make([]int, 4, 8)
	fmt.Print(s2, "\n")

	test(s2)
	fmt.Print(s2, "\n")

	// string 是不可变的
	var str string
	str = "abcdefg"
	s3 := str[:]
	fmt.Print(s3, "\n")
	// string  不能直接修改，如果要修改需要转换为数组，再转回去赋值
	arr1 := []byte(str)
	arr1[0] = 'z'
	str = string(arr1)
	fmt.Print(str, "\n")
	// 如果有中文 需要使用 rune
	arr2 := []rune(str)
	arr2[0] = '好'
	str = string(arr2)
	fmt.Print(str, "\n")

	tab()
	//paixu(5)
	//chazhao()

}
