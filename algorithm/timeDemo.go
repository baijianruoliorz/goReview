package algorithm

import (
	"fmt"
	"time"
)

/*
*  @author liqiqiorz
*  @data 2020/11/7 22:01
 */
func TimeDemo(){
	now:=time.Now()
	fmt.Println("current time:%v\n",now)
	year:=now.Year()
	month:=now.Month()
	day:=now.Day()
	hour:=now.Hour()
	minute:=now.Minute()
	second:=now.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}
