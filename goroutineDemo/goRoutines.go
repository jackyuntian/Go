package goroutineDemo

import (
	"fmt"
	"sync"
)

/**
    @date: 2021/12/7
**/

var Wg sync.WaitGroup

func Hellos(i int){
	defer Wg.Done() // goroutine 结束就登记-1
	fmt.Println("Hello Goroutine!", i)
}
