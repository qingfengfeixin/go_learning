package remote_package

// 引入之前可以使用go get -u github.com/easierway/concurrent_map 进行代码下载
import (
	"testing"

	cm "github.com/easierway/concurrent_map"
)

func TestConcurrentMap(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
}
