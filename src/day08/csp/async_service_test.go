package csp

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 500)
	return "Done"
}

func otherTask() {
	fmt.Println("working on sth else...")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

//TestService 串行执行
func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func AsyncService() chan string {
	//retCh := make(chan string)    //声明一个channel
	retCh := make(chan string, 1) //声明一个buffer为1 的channel
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret // 往channel里面放数据
		fmt.Println("service exited.")
	}()
	return retCh
}

func TestAsynService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh) // 从channel里面取数据
}

func TestSelect(t *testing.T) {
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 100):
		t.Error("time out")
	}
}
