package test

import (
	"fmt"
	"testing"
)

func TestFoo(t *testing.T) {
	if 1 != 1 {
		t.Fail()
	}
	fmt.Println("Test success")
}

func BenchmarkFoo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Printf("hello")
	}
}

func Example() {
	if true {
		fmt.Println("hello world")
	} else {
		fmt.Println("wrong")
	}
	// Output:
	// hello world
}

func swap0(x, y int) {
	var temp int
	temp = x
	x = y
	y = temp
}

func Swap1(x, y int) (int, int) {
	return y, x
}

func swap2(x, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

func ExampleSwap1() {
	var x, y = 3, 9
	fmt.Println(x, y)
	swap0(x, y)
	fmt.Println(x, y)
	x, y = Swap1(x, y)
	fmt.Println(x, y)
	swap2(&x, &y)
	fmt.Println(x, y)
	//Output:
	//3 9
	//3 9
	//9 3
	//3 9
}

type intclass struct {
	value int
}

func SwapIntclass(x, y intclass) {
	var temp int
	temp = x.value
	x.value = y.value
	y.value = temp
}

func SwapIntclass2(x, y *intclass) {
	var temp int
	temp = x.value
	x.value = y.value
	y.value = temp
}

func ExampleSwapIntclass() {
	var x, y = intclass{3}, intclass{9}
	fmt.Println(x.value, y.value)
	SwapIntclass(x, y)
	fmt.Println(x.value, y.value)
	SwapIntclass2(&x, &y)
	fmt.Println(x.value, y.value)
	//Output:
	//3 9
	//3 9
	//9 3
}
