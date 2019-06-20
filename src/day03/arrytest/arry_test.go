package arrytest

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr3 := [...]int{1, 2, 3, 4, 5, 6, 7} //可以不指定长度，根据后面的赋值决定长度
	arr1[1] = 15
	t.Log(arr1[1], arr1)
	t.Log(arr, len(arr3))
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5, 6}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}
	t.Log(arr3[1:3]) //数组的切片,左闭右开
	t.Log(arr3[1:])
	t.Log(arr3[:4])
}
