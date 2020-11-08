package main

import (
	"algorithm/reflectDemo"
	"fmt"
)

/*
*  @author liqiqiorz
*  @data 2020/11/7 22:09
 */
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
func main() {
	//algorithm2.TimeDemo()

	//[1 2 3]
	//[1 3]
	//[1 3 3]
	//type Map map[string][]int
	//m:=make(Map)
	//s:=[]int{1,2}
	//s=append(s,3)
	//fmt.Printf("%+v\n",s)
	//m["qimi"]=s
	//s=append(s[:1],s[2:]...)
	//fmt.Printf("%+v\n",s)
	//fmt.Printf("%+v\n",m["qimi"])

	//A 1 2 3
	//B 10 2 12
	//BB 10 12 22
	//AA 1 3 4
	x := 1
	y := 2
	//defer注册要延迟执行的函数时该函数所有的参数都需要确定其值
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20

	//func A
	//recover in B
	//func C

	//algorithm.FuncA()
	//algorithm.FuncB()
	//algorithm.FuncC()

	//algorithm.ReadFile("../algorithm/IOkeyWord.go")
	var a float32 = 3.14
	reflectDemo.ReflectType(a)
	var b = 100
	reflectDemo.ReflectType(b)
}
