package emptyinterface

import (
	"fmt"
	"testing"
)

func Dosth(p interface{}) {
	if i, ok := p.(int); ok {
		fmt.Println("Interger", i)
		return
	}
	if s, ok := p.(string); ok {
		fmt.Println("String", s)
		return
	}
	fmt.Println("unknown type")
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	Dosth(10)
	Dosth("10")
}

func Dosth1(p interface{}) {
	switch v := p.(type) {
	case int:
		fmt.Println("Interger", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("unknown type")
	}
}

func TestEmptyInterfaceAssertion1(t *testing.T) {
	Dosth1(10)
	Dosth1("10")
}
