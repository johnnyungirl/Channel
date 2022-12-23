package main

import (
	"log"
	"math/rand"
	"time"
)

func longTimeRequest() <-chan int32 {
	r := make(chan int32)
	go func() {
		time.Sleep(3 * time.Second)
		r <- rand.Int31n(100)
	}()
	return r
}
func SumQuares(a, b int32) int32 {
	return a*a + b*b
}
func main() {
	rand.Seed(time.Now().Unix())
	a, b := longTimeRequest(), longTimeRequest()
	log.Println(SumQuares(<-a, <-b))

}
