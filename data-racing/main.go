package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var count int = 0

	lock := new(sync.RWMutex)

	for i := 1; i <= 5; i++ {
		go func() {
			for j := 1; j <= 10000; j++ {
				lock.Lock()
				count += 1
				lock.Unlock()
			}
		}()
	}

	time.Sleep(time.Second * 4)
	fmt.Println(count)

}
