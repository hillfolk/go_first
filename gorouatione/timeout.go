//select 문을 이용해서 타임아웃 기능을 쉽게 구현 할수 있다.


package main

import (
	"fmt"
	"time"
)

func main() {
	quit := make(chan struct{})
	done := process(quit)
	timeout := time.After(3 * time.Second)
	select {
	case d := <- done:
		fmt.Println(d)
	case <-timeout: //일정 시간후 체널 데이터 발생 
		fmt.Println("Time out!")
		quit <- struct{}{}
	}
}


func process(quit <-chan struct{}) chan string {
	done := make(chan string)

	go func() {
		go func() {
			fmt.Println("Start")
			c := time.Tick(1 * time.Second)
			for now := range c {
				fmt.Printf("%v \n", now)
			}
			time.Sleep(10 * time.Second)
			fmt.Println("End")
			done <- "Complete!"
			
		}()
	<- quit
	return
}()

return done
	
}
