package main

import (
	"fmt"
	"net"
	"sync"
)

//使用goroutine实现并发端口扫描
func main() {

	//A WaitGroup waits for a collection of goroutines to finish.
	var wg sync.WaitGroup

	for i:=1;i<65335;i++ {

			//内部计数器 递增+1
			wg.Add(1)
			go func(j int) {

				//完成后将计数器减少1
				defer wg.Done()
				adress :=fmt.Sprintf("127.0.0.1:%d",j)
				conn,err:=net.Dial("tcp",adress)
				if err !=nil{
					return
				}
				//断开连接
				conn.Close()
				fmt.Printf("%d is opening\n",j)

			}(i)
	}
	//直到内部计数器达到零，再执行下一步
	wg.Wait()
}
