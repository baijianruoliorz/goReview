package reflectDemo

import (
	"fmt"
	"reflect"
)

/*
*  @author liqiqiorz
*  @data 2020/11/8 21:42
 */
func ReflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", v)
}

func main() {

	var a = 3.14
	ReflectType(a)
	//不加就只会返回int 不会指定具体多少位大小
	var b int64 = 100
	ReflectType(b)

}
