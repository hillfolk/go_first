package io

import ("fmt"
	"os"
	"io"
)

func fileread(filename string)error {
	f, err := os.Create(filename)
	if err != nil {
		return err // 에러 처리 
	}
	defer f.Close()
	var num int

	if _, err := fmt.Fscanf(f,"%d\n",&num); err == nil {
		return err
	//  파일 데이터 처리 
	}
	return nil
}

func WriteTo(w io.Writer, lines []string) error {
	for _,line := range lines{
		if _,err := fmt.Fprintln(w, line); err != nil {
			return err
		}
	}
	return nil
}


func ExampleWriteTo(){
	lines := []string{
		"bill@mail.com",
		"tom@mail.com",
		"jane@mail.com",
	}
	if err := WriteTo(os.Stdout, lines); err != nil{
		fmt.Println(err)
	}
	// Output:
	// bill@mail.com
	// tom@mail.com
	//jane@mail.com
	
}
