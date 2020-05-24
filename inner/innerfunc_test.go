package inner

import (
	"fmt"
	"testing"
)

func ExamplestudyMake() {
	studyMake()
	// Output:
	// [0 0 0 0 0]
	// [0 0 0 0 0]
	// [1 2]
}

func TestSqrt(t *testing.T) {
	if sqrt(16) != 4 {
		t.Fail()
	}
}

func updateArray(arr [5]int) {
	if len(arr) > 1 {
		arr[1] = 99
	}
}

func updateSlice(slice []int) {
	if len(slice) > 1 {
		slice[1] = 99
	}
}

func ExampleSum() {
	list1 := []float64 {2.3, 4.5, 1.2, 4.5, 5.6}
	fmt.Printf("sum is %f\n", Sum(list1))
	// Output:
	// sum is 18.100000
}


func Example() {
	var arr1 = [...]int{1, 2, 3, 4, 5}
	fmt.Printf("len is %d ,cap is %d,  slice is %v\n", len(arr1), cap(arr1), arr1)

	sl1 := arr1[1:4]
	fmt.Printf("len is %d ,cap is %d,  slice is %v\n", len(sl1), cap(sl1), sl1)

	sl2 := make([]int, 6, 60)
	fmt.Printf("len is %d ,cap is %d,  slice is %v\n", len(sl2), cap(sl2), sl2)

	updateArray(arr1)
	fmt.Printf("len is %d ,cap is %d,  slice is %v\n", len(arr1), cap(arr1), arr1)

	updateSlice(sl1)
	fmt.Printf("len is %d ,cap is %d,  slice is %v\n", len(sl1), cap(sl1), sl1)
	
	sl1 = append(sl1, 111)
	sl1 = append(sl1, 222)
	sl1 = append(sl1, 333)
	fmt.Printf("len is %d ,cap is %d,  slice is %v\n", len(sl1), cap(sl1), sl1)
	// Output:
	// len is 5 ,cap is 5,  slice is [1 2 3 4 5]
	// len is 3 ,cap is 4,  slice is [2 3 4]
	// len is 6 ,cap is 60,  slice is [0 0 0 0 0 0]
	// len is 5 ,cap is 5,  slice is [1 2 3 4 5]
	// len is 3 ,cap is 4,  slice is [2 99 4]
	// len is 6 ,cap is 8,  slice is [2 99 4 111 222 333]
}
