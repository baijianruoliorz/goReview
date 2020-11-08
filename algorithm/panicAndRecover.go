package algorithm

import "fmt"

/*
*  @author liqiqiorz
*  @data 2020/11/8 15:20
 */

func FuncA(){
	fmt.Println("func A")
}
func FuncB(){
	defer func() {
		err:=recover()
		if err!=nil{
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")
}
func FuncC(){
	fmt.Println("func C")
}
