package main

import "fmt"
import "os"
import "strconv"

func factorial(n uint64) uint64 {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func main() {
	args0,err := strconv.ParseUint(os.Args[1],10,64)
	fmt.Println(err)
	fmt.Println(args0)
	fmt.Println(factorial(args0)) // 120
		
}
