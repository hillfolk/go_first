package string_go

import ("fmt"
	"strconv")

func Example_strtonumber() {
	var i int
	var k int64
	var f float64
	var s string
	var err error

	i, err = strconv.Atoi("350")                  // i == 350
	k, err = strconv.ParseInt("cc7fdd", 16, 32) // k == 13402077
	k, err = strconv.ParseInt("0xcc7fdd", 0, 32)  // k == 13402077
	f, err = strconv.ParseFloat("3.14", 64)    // f == 3.14
	s = strconv.Itoa(340)                     // s == "340"
	s = strconv.FormatInt(13402077, 16)      // s == "cc7fdd"
	
	fmt.Println(i)
	fmt.Println(k)
	fmt.Println(f)
	fmt.Println(s)
	fmt.Println(err)
	
	// Output:
	// .
}
