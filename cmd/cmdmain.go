package main

import (
	"encoding/json"
	"fmt"
	"study/inner"
	"study/structype"
)

func say(s string, ch chan string) {
	for i := 0; i < 5; i++ {
		//time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
		ch <- s
	}
	close(ch)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

func main() {
	ch := make(chan string, 3)
	go say("say", ch)
	stream := ""
	for rcv := range ch {
		stream += rcv
		fmt.Println("receive :" + rcv)
	}
	fmt.Println("receive all :" + stream)
	fmt.Println("hell World")
	usr1 := structype.User{
		Name: "nick",
		Age:  30,
	}
	conJSON, _ := json.Marshal(usr1)
	fmt.Println(string(conJSON)) //{"userName":"nick","userAge":0}

	list1 := []float64{2.3, 4.5, 1.2, 4.5, 5.6}
	fmt.Printf("sum is %f\n", inner.Sum(list1))
	fmt.Printf("Hello %s\n", "world")

	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)

}
