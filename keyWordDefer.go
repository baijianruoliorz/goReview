package algorithm

import "fmt"

/*
*  @author liqiqiorz
*  @data 2020/11/8 14:59
 */
func calc(index string, a,b int)int{
	ret:=a+b
	fmt.Println(index,a,b,ret)
	return ret
}
func main(){
	x:=1
	y:=2
	//defer注册要延迟执行的函数时该函数所有的参数都需要确定其值
	defer  calc("AA",x,calc("A",x,y))
	x=10
	defer calc("BB",x,calc("B",x,y))
	y=20

}