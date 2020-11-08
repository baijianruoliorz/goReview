package algorithm

import (
	"fmt"
	"io/ioutil"
)

/*
*  @author liqiqiorz
*  @data 2020/11/8 21:22
 */
//ioutil的文件读取操作
func ioutileReadFile(name string) {
	file, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println("read file failed,err", err)
		return
	}
	fmt.Println(string(file))
}
