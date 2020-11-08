package algorithm

import (
	"fmt"
	"io/ioutil"
)

/*
*  @author liqiqiorz
*  @data 2020/11/8 21:33
 */
func ioutilWriteFile(name string) {
	str := "hello baijianruoliorz"
	err := ioutil.WriteFile("./xx,txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed,err", err)
		return
	}
}
