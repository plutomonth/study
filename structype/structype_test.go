package structype

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircle(t *testing.T) {
	c := circle{
		radius: 3,
	}
	if math.Abs(c.Area()-3.1415*c.radius*c.radius) > 0.01 {
		t.Fail()
	}

	if math.Abs(c.Perimeter()-3.1415*c.radius*2) > 0.01 {
		t.Fail()
	}
}

func TestFlat(t *testing.T) {
	f := flat{
		8,
		6,
	}
	if math.Abs(f.Area()-f.Heigth*f.Width) > 0.01 {
		t.Fail()
	}

	if math.Abs(f.Perimeter()-2*(f.Heigth+f.Width)) > 0.01 {
		t.Fail()
	}
}

func TestManager(t *testing.T) {
	user := User{
		"John",
		12,
	}

	manager := Manager{
		&user,
	}

	admin := Admin{
		user,
	}

	if manager.Age != 12 {
		t.Fail()
	}

	if admin.Age != 12 {
		t.Fail()
	}
}

func Example() {
	f := flat{
		Heigth: 8,
		Width:  6,
	}
	fmt.Printf(f.ToJSON())
	// Output:
	// {"heigth":8,"width":6}

}

func TestAdd(t *testing.T) {
	var a Integer = 10
	println("Firt time, return value is ", a.Add(Integer(10)))
	assert.Equal(t, a, Integer(20))

	println("Second time, return value is ", a.AddByValue(Integer(10)))
	assert.Equal(t, a, Integer(20))
}
