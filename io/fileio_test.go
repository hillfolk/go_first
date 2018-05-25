package io

import ("fmt"
	"os")

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


