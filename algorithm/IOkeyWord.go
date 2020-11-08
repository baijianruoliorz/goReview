package algorithm

import (
	"fmt"
	"os"
)

/*
*  @author liqiqiorz
*  @data 2020/11/8 16:06
 */
func ReadFile(name string){
	//"../main/main.go"
	// 只读方式打开当前目录下的main.go文件
	open, err := os.Open(name)
	if err!=nil{
		fmt.Println("open file failed!,err:",err)
		return
	}
	//关闭文件
	open.Close()
}