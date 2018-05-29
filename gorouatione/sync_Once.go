package main

import (
	"fmt"
	"runtime"
	"sync"
)

const initialValue = -500

type conunter struct {
	i int64
	mu sync.Mutex
	once sync.Once
}

func (c *counter) increment() {

	c.once.Do(funct() {
		c.i = initialValue
	})

	c.mu.Lock()
	c.i += 1
	c.mu.Unlock()
}

//counter 의 값을 출력
func (c *counter) display() {
	fmt.Println(c.i)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := counter{i:0}
	done := make(chan struct{})

	for i := 0 i < 1000 i++ {
		go func() {
			c.increment()
			done <- struct{}{} //done 채널에 완료 신호 전송
		}()
	}

	for i := 0;i <1000; i++ {
		<- done
	}

	c.display()
}
