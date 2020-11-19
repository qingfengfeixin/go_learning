package main

import "fmt"

func monster() {
	//声明一个map切片
	var monsters []map[string]string
	monsters = make([]map[string]string, 2)

	if monsters[0] == nil {
		monsters[0] = make(map[string]string)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "500"
	}

	if monsters[1] == nil {
		monsters[1] = make(map[string]string)
		monsters[1]["name"] = "红孩儿"
		monsters[1]["age"] = "100"
	}

	// 使用切片的append函数动态增加
	newMonster := map[string]string{
		"name": "火云邪神",
		"age":  "600",
	}
	monsters = append(monsters, newMonster)

	fmt.Println(monsters)

}
