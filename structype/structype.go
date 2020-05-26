package structype

import (
	"encoding/json"
	"math"
)

// Sharp type
type Sharp interface {
	Area() float64
	Perimeter() float64
}

type circle struct {
	radius float64
}

func (c *circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// Flat type
type flat struct {
	Heigth float64 `json:"heigth"`
	Width  float64 `json:"width"`
}

// Area method of
func (f *flat) Area() float64 {
	return f.Heigth * f.Width
}

func (f *flat) Perimeter() float64 {
	return 2 * (f.Heigth + f.Width)
}

func (f *flat) ToJSON() string {
	conJSON, _ := json.Marshal(f)
	return string(conJSON)
}

// User type
type User struct {
	Name string `json:"userName"`
	Age  int    `json:"userAge"`
}

// Manager type
type Manager struct {
	*User
}

// Admin type
type Admin struct {
	User
}
