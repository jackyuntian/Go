package goroutineDemo

import (
	"fmt"
)

/**
    @date: 2021/12/7
**/

func channleDemo (){
	// 声明一个通道 引用类型为空
	var ch chan int
	// 开辟一个通道
	ch =make(chan int)
	go recv(ch)
	// 将一个值发送到通道中
	ch <- 10
	// 从一个通道中接收值

	//x := <- ch // 从ch中接收值并赋值给变量x
	defer close(ch)
	//<-ch   // 从ch中接收值,忽略结果
	fmt.Println("发送成功")
}


func recv (c chan int) {
	ret := <-c
	fmt.Println("接收成功",ret)
}