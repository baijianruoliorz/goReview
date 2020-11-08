package main

import "fmt"

/*
*  @author liqiqiorz
*  @data 2020/11/8 14:47
 */
func main(){
	type Map map[string][]int
	m:=make(Map)
	s:=[]int{1,2}
	s=append(s,3)
	fmt.Printf("%+v\n",s)
	m["qimi"]=s
	s=append(s[:1],s[2:]...)
	fmt.Printf("%+v\n",s)
	fmt.Printf("%+v\n",m["qimi"])


}
