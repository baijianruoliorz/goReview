package algorithm

import "strings"

/*
*  @author liqiqiorz
*  @data 2020/11/7 21:31
 */
//一句话中统计字母个数
func WordCount(words string) (r map[string]int){
	wordsArr:=strings.Split(words,"")
	r=make(map[string]int,len(wordsArr))
	for _,word:=range wordsArr{
		r[word]++
	}
	return
}

