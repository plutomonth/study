package inner

import (
	"errors"
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

func sqrt(value float64) (float64, error) {
	if value < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	return math.Sqrt(value), nil
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
	for i := 0; i < size; i++ {
		sum += list[i]
	}
	return sum
}

// SumByRange method
func SumByRange(list []float64) float64 {
	sum := 0.0
	for _, v := range list {
		sum += v
	}
	return sum
}

// Capitals method
func Capitals() map[string]string {
	capitalMap := make(map[string]string)
	capitalMap["France"] = "Paris"
	capitalMap["Italy"] = "Rome"
	capitalMap["Japan"] = "Toyko"
	return capitalMap
}

// Less , the anonymous func 
var Less = func(a, b int, z float64) bool {
	return a*b <int(z)
}

