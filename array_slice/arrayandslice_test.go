package arrayandslice

import "fmt"

func Example_array() {
	fruits := [3]string{"사과","바나나","토마토"}
	for _, fruit := range fruits {
		fmt.Printf("%s는 맛있다.\n", fruit)	
	}
	// Output:
	// 사과는 맛있다.
	// 바나나는 맛있다.
	// 토마토는 맛있다.
}

func Example_slice() {

	fruits := make([]string, 3)
	fruits[0] = "사과"
	fruits[1] = "바나나"
	fruits[2] = "토마토"
	
	for _, fruit := range fruits {
		fmt.Printf("%s는 맛있다.\n", fruit)	
	}
	// Output:
	// 사과는 맛있다.
	// 바나나는 맛있다.
	// 토마토는 맛있다.
	
}

func Example_sling() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	fmt.Println(nums[0:3])
	fmt.Println(nums[1:])
	fmt.Println(nums[:3])
	//Output:
	//.
}

	
