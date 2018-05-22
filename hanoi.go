package main

import "fmt"

var count int

func Move(n, from, to, via int){
     count++
     fmt.Println("count:",count)
     if n <= 0 {
     return
     }
     Move(n-1,from,via, to)
     fmt.Println(from, "->", to)
     Move(n-1,via,to, from)
    }		      

func Hanoi(n int){
     fmt.Println("Number of Disk:",n)
     Move(n, 1,2,3)
     
     }

func main() {
     Hanoi(100)
}