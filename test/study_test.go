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
		fmt.Sprintf("hello")
	}
}

func ExampleSqrt() {
	if true {
		fmt.Println("hello world")
	} else {
		fmt.Println("wrong")
	}
	// Output:
	// hello world
}
