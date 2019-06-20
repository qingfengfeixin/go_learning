package sharemem

import (
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)   // 手动等待协程完成，输出counter 防止test跑完，协程没有跑完导致结果不准确
	t.Logf("counter=%d", counter) //结果小于5000，因为线程不安全，数据不一致，导致数据丢失操作
}

func TestCounterThreadSafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter=%d", counter) // 通过锁，获得5000
}

func TestCounterWaitGroup(t *testing.T) {
	var (
		mut sync.Mutex
		wg  sync.WaitGroup
	)

	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1) //
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done() //
		}()
	}
	//time.Sleep(1 * time.Second)
	wg.Wait() //等待每个协程完成
	t.Logf("counter=%d", counter)
}
