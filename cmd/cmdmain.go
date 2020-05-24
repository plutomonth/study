package main

import (
	"encoding/json"
	"fmt"
	"study/structype"
	"study/inner"
)

func main() {
	fmt.Println("hell World")
	usr1 := structype.User{
		Name: "nick",
		Age:  30,
	}
	conJSON, _ := json.Marshal(usr1)
	fmt.Println(string(conJSON)) //{"userName":"nick","userAge":0}

	list1 := []float64 {2.3, 4.5, 1.2, 4.5, 5.6}
	fmt.Printf("sum is %f\n", inner.Sum(list1))
	fmt.Printf("Hello %s\n", "world")
}
