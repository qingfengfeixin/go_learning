package main

// defer panic recover

import (
	"fmt"
)

func test() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	num1 := 1
	num2 := 0
	num3 := num1 / num2
	fmt.Print(num3)
}

func main() {
	test()
	fmt.Print("recover from panic \n")

}
