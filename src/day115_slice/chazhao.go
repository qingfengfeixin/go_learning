package main

import "fmt"

func shunxuchazhao() {
	str := [4]string{"白眉鹰王", "金毛狮王", "紫衫龙王", "青翼蝠王"}

	var heroName string
	fmt.Scanln(&heroName)
	for i, v := range str {
		if v == heroName {
			fmt.Print("find ", heroName, "下标为", i, "\n")
			break
		}
		if i == len(str)-1 {
			fmt.Print("not find! ", heroName, "\n")
		}
	}
}

func erfenchazhao(arr []int, left int, right int, num int) {

	if left > right {
		fmt.Println("not find ", num)
		return
	}

	mid := (left + right) / 2

	if arr[mid] < num {
		erfenchazhao(arr, mid+1, right, num)
	} else if arr[mid] > num {
		erfenchazhao(arr, left, mid-1, num)
	} else {
		fmt.Print("find ", num, " index is ", mid, "\n")
	}

}

func chazhao() {
	//shunxuchazhao()

	s1 := data(8)
	maopao(s1)

	var num int
	fmt.Println("请输入要查找的数值：")
	fmt.Scanln(&num)

	erfenchazhao(s1, 0, len(s1)-1, num)
}
