package reflectDemo

import (
	"reflect"
)

/*
*  @author liqiqiorz
*  @data 2020/11/8 22:00
 */

func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {

		//修改的是副本，reflect包会引发panic
		v.SetInt(200)
	}
}
func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	// 反射中使用 Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}
