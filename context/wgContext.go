package context

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
*  @author liqiqiorz
*  @data 2020/11/9 11:25
 */
var wg sync.WaitGroup

func Worker(ctx context.Context) {
LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // 等待上级通知
			break LOOP
		default:
		}
	}
	wg.Done()
}

func main() {
	//wg.Add(1)
	//go Worker()
	//wg.Wait()
	//fmt.Println("over")
}
