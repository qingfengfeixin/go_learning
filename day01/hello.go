package main

// helloworld
import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println("hello ", os.Args[1])
	}
	fmt.Println("Hello, World!")
}
