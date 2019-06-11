package loop

// go中的循环只用for 没有while

import "testing"

func TestWhicleLoop(t *testing.T) {
	n := 0
	// while(n<5)
	for n < 5 {
		t.Log(n)
		n++
	}
	// 无限循环
	// while(true)
	n = 0
	for {
		t.Log(n)
		n++
		if n > 10 {
			break
		}
	}
}

func TestIfMultiSec(t *testing.T) {
	if a := 1 == 1; a {
		t.Log("1==1")
	}
}

func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log(i, "偶数")
		case 1, 3:
			t.Log(i, "奇数")
		default:
			t.Log(i, "not 0-3")
		}
	}
}

func TestSwitchCaseCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log(i, "Even")
		case i%2 == 1:
			t.Log(i, "Odd")
		default:
			t.Log(i, "unknown")
		}
	}
}
