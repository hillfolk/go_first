package string_go

import "fmt"

func Example_str() {
	s := "abc"
	ps := &s
	s += "def"
	fmt.Println(s)
	fmt.Println(*ps)

	// Output:
	// abcdef
	// abcde
}
