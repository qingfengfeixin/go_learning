package series

import "fmt"

//定义init
func init() {
	fmt.Println("init1")
}

// 也可以定义两个init
func init() {
	fmt.Println("init2")
}

// GetFibSerie 首字母大写代表可以被包外的程序访问,如果是小写，则包外无法访问
func GetFibSerie(n int) []int {
	ret := []int{1, 2}
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-2]+ret[i-1])
	}
	return ret
}
