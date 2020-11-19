package main

import "fmt"

type Point struct {
	x, y int
}

type Rect struct {
	leftUp, rightDown Point
}

func point() {
	r1 := Rect{Point{1, 2}, Point{3, 4}}
	fmt.Println(&r1.leftUp.x, &r1.leftUp.y, &r1.rightDown.x, &r1.rightDown.y)
}
