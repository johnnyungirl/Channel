package main

import (
	"fmt"
	"math/rand"
	"time"
)

func longTimeRequest(r chan<- int32) {
	time.Sleep(3 * time.Second)
	r <- rand.Int31n(100)
}
func SumQuares(a int32, b int32) int32 {
	return a*a + b*b
}
func main() {
	rand.Seed(time.Now().Unix())
	result:=make(chan int32)
	go longTimeRequest(result)
	go longTimeRequest(result)
	fmt.Println(SumQuares(<-result,<-result))
}
