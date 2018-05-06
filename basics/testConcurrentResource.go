package main

import (
	"sync"
	"runtime"
	"fmt"
	"sync/atomic"
)

var (
	count int32
	wg    sync.WaitGroup
	mutex sync.Mutex
)

func addCountAtomic() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		value := atomic.LoadInt32(&count)
		runtime.Gosched()
		value++
		atomic.StoreInt32(&count, value)
	}
}

func addCountLock() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		mutex.Lock()
		value := count
		runtime.Gosched()
		value++
		count = value
		mutex.Unlock()
	}
}

func main() {
	wg.Add(2)
	go addCountLock()
	go addCountLock()
	wg.Wait()
	fmt.Print(count)
}
