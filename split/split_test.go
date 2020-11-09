package split

import (
	"reflect"
	"testing"
)

/*
*  @author liqiqiorz
*  @data 2020/11/8 22:13
 */
//cd 到这个文件夹
//PASS
//ok      algorithm/split 0.191s

func TestSplit(t *testing.T) {
	// 测试函数名必须以Test开头，必须接收一个*testing.T类型参数
	got := Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("execept:%v,got:%v", want, got)
	}
}
