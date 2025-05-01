// mutex and rwmutex
package main

import (
	"fmt"
	"sync"
)

func main() {
	cnt := 1
	mux := *&sync.RWMutex{}
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			mux.RLock()
			fmt.Println(cnt, "--", i)
			mux.RUnlock()
			defer wg.Done()
			mux.Lock()
			defer mux.Unlock()
			cnt *= 5
		}()
	}
	wg.Wait()
	fmt.Println(cnt)

}
