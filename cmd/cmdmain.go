package main

import (
	"encoding/json"
	"fmt"
	"study/structype"
)

func main() {
	fmt.Println("hell World")
	usr1 := structype.User{
		Name: "nick",
		Age:  30,
	}
	conJSON, _ := json.Marshal(usr1)
	fmt.Println(string(conJSON)) //{"userName":"nick","userAge":0}
}
