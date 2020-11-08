package algorithm

import (
	"bufio"
	"fmt"
	"os"
)

/*
*  @author liqiqiorz
*  @data 2020/11/8 21:29
 */
func Bufio(name string) {
	file, err := os.OpenFile("xx.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed,err: ", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		//数据写入缓存
		writer.WriteString("hello baijianruoliorz\n")
	}
	//缓存中的数据写入文件
	writer.Flush()

}
