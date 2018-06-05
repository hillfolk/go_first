package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {
	client, err := net.Dial("tcp","127.0.0.1:8000")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer client.Close()

	go func(c net.Conn) {
		data := make([]byte, 4096)

		for {
			n, err := c.Read(data) // 서버에서 받은 데이터를 읽음

			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(string(data[:n])) //데이터 출력
			time.Sleep(1 * time.Millisecond)
			
		}
		
	}(client)

	go func(c net.Conn) {
		i := 0

		for {
			s := "Hello " + strconv.Itoa(i)

			_, err := c.Write([]byte(s)) //서버 데이터를 보냄
			if err != nil {
				fmt.Println(err)
				return
			}

			i++
			time.Sleep(1 * time.Millisecond)	
		}
	}(client)

	fmt.Scanln()
}
