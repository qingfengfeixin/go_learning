package main

import "fmt"

func calc(args ...int) (sum int) {
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return
}

func main() {
	res := calc(1, 2, 3, 4, 6)
	fmt.Println("res = ", res)

	a := 1
	b := 2
	a, b = b, a
	fmt.Print("a = ", a, "b = ", b, "\n")

	num1 := 100
	fmt.Println(num1, &num1, "\n")

	num2 := new(int)
	*num2 = 100
	fmt.Println(num2, &num2, *num2, "\n")

}
