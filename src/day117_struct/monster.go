package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name  string `json:"name"` // struct tag
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

// 方法是值复制
func (m Monster) callName() {
	(m).Name = "铁扇公主"
	fmt.Println(m.Name)
}

//传入地址改变原结构体属性
func (m *Monster) callName1() {
	(*m).Name = "铁扇公主"
}

func (m *Monster) String() string {
	str := fmt.Sprintf("Name =[%v], Age = [%v]", m.Name, m.Age)
	return str
}

func monster() {
	mons1 := Monster{"牛魔王", 500, "芭蕉扇"}
	mons1.callName()

	if jsonStr, err := json.Marshal(mons1); err != nil {
		fmt.Println("json 处理错误", err)
	} else {
		fmt.Println("jsonStr = ", jsonStr, "\n", string(jsonStr))
	}
	(&mons1).callName1()
	fmt.Println(mons1.Name)
	fmt.Println(mons1)
	// 如果实现了  *Monster 类型的 String方法，就会自动调用
	fmt.Println(&mons1)
}
