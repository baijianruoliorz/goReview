package algorithm

import (
	"fmt"
	"os"
)

/*
*  @author liqiqiorz
*  @data 2020/11/8 21:26
 */
func ioutilWrite() {
	file, err := os.OpenFile("xx.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed,err", err)
		return
	}
	defer file.Close()
	//写字节切片
	str := "hello baijianruoliorz"
	file.Write([]byte(str))
	//直接写字符串
	file.WriteString("hello liqiqiorz")

}
