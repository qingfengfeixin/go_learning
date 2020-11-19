package main

import (
	"fmt"
)

func intersection(nums1 []int, nums2 []int) []int {
	// 使用两个map 如果在其中 则输出

	var result []int = make([]int, 0)
	var m1 map[int]int = make(map[int]int)
	var m2 map[int]int = make(map[int]int)

	for _, v := range nums1 {
		m1[v] = 1
	}

	for _, v := range nums2 {
		m2[v] = 2
	}

	for i, _ := range m1 {
		if _, ok := m2[i]; ok {
			result = append(result, i)
		}

	}

	return result

}

func findContinuousSequence(target int) (res [][]int) {
	// 窗口滑动
	res = make([][]int, 0)

	var t int = 0
	for i := 1; i < target; i++ {
		t = i
		var r1 []int = make([]int, 0)
		r1 = append(r1, i)
		for j := i + 1; j < target; j++ {
			t += j
			r1 = append(r1, j)
			if t == target {
				res = append(res, r1)
			}

		}

	}
	return

}

func removeDuplicates(nums []int) int {
	var m1 map[int]int = make(map[int]int)
	var s1 []int
	for i := 0; i < len(nums); i++ {

		if _, ok := m1[nums[i]]; !ok {
			m1[nums[i]] = 1
			s1 = append(s1, nums[i])
		}
	}
	nums = s1
	fmt.Println(s1)
	fmt.Println(nums)
	return len(s1)

}

func singleNumber(nums []int) (v_o int) {
	var m1 map[int]int = make(map[int]int)
	for _, v := range nums {
		if _, ok := m1[v]; ok {
			m1[v] += 1
		} else {
			m1[v] = 1
		}

	}
	for v := range m1 {
		if m1[v] == 1 {
			v_o = v
		}
	}
	return v_o
}
func main() {
	s1 := []int{1, 2, 2, 3, 5, 5}
	s2 := []int{2, 3, 10}
	s3 := intersection(s1, s2)
	fmt.Println(s3)
	s4 := []int{4, 1, 2, 1, 2}
	findContinuousSequence(15)
	removeDuplicates(s1)
	fmt.Println(s1)
	fmt.Println(singleNumber(s4))

}
