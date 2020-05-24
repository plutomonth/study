package inner

import (
	"fmt"
	"math"
)

func studyMake() {
	var slice1 []int = make([]int, 5, 10)
	fmt.Println(slice1)

	var slice2 []int = make([]int, 5)
	fmt.Println(slice2)

	var slice3 []int = []int{1, 2}
	fmt.Println(slice3)
}

func sqrt(value float64) float64 {
	return math.Sqrt(value)
}

func getSequence(step int) func() int {
	i := 0
	return func() int {
		i += step
		return i
	}
}

// Sum method
func Sum(list []float64) float64 {
	size := len(list)
	sum := 0.0
	for i:= 0; i<size; i++ {
		sum += list[i]
	}
	return sum
}