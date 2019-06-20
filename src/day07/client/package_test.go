package client

// 注意此处import是需要在go_path环境变量下的src目录下
import (
	"day07/series"
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibSerie(5))
}
