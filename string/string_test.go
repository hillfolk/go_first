package string_go

import "fmt"

func Example_strCat() {
	s := "abc"
	ps := &s
	s += "def"
	fmt.Println(s)
	fmt.Println(*ps)

	// Output:
	// abcdef
	// abcdef
}

func Example_pirntBytes2(){
	s := "가나다"
	fmt.Printf("%x\n",s)
	fmt.Printf("% x\n",s)
	// Output:
	// eab080eb8298eb8ba4
	// ea b0 80 eb 82 98 eb 8b a4
	
}
	
	
