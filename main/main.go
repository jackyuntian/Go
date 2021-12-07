package main

import "yuntianSuiBi/Go/goroutineDemo"

/**
    @date: 2021/12/7
**/
func main(){
	// 关键字go 开启goroutine
	//go goroutineDemo.Hello()
	//fmt.Println("打印")
	//time.Sleep(time.Second)
	for i:=0; i<10; i++ {
		goroutineDemo.Wg.Add(1)
		go goroutineDemo.Hellos(i)
	}
	goroutineDemo.Wg.Wait()
}

