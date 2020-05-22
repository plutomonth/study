package inner

import "testing"

func ExampleStudymake() {
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
