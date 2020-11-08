package algorithm

import (
	"fmt"
	"io"
	"os"
)

/*
*  @author liqiqiorz
*  @data 2020/11/8 16:20
 */
func readFile(name string){
	file, err := os.Open(name)
	if err!=nil{
		fmt.Println("open file failed!err:",err)
		return
	}
	defer file.Close()
//	使用read方法读取数据
    var tmp=make([]byte,128)
    n,err:=file.Read(tmp)
    if err==io.EOF{
    	fmt.Println("文件读完了")
		return
	}
	if err!=nil{
		fmt.Println("read file failed,err:",err)
		return
	}
	fmt.Printf("读取了%d字节数据\n",n)
    fmt.Println(string(tmp[:n]))


}