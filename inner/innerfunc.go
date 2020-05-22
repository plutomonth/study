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

func main() {
	studyMake()
	fmt.Println("math.Sqrt(17)  =", math.Sqrt(17))

	varSqrtFunc := func(value float64) float64 {
		return math.Sqrt(value)
	}
	fmt.Println("varSqrtFunc(17)=", varSqrtFunc(17))

	nextValue := getSequence(2)
	fmt.Println(nextValue())
	fmt.Println(nextValue())
	fmt.Println(getSequence(3)())
	fmt.Println(getSequence(4)())
}
