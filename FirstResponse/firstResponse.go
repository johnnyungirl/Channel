package main

import (
	"fmt"
	"math/rand"
	"time"
)

func source(c chan<- int32) {
	ra := rand.Int31()
	rb := rand.Int31n(3) + 1
	time.Sleep(time.Second * time.Duration(rb))
	c <- ra
}
func main() {
	rand.Seed(time.Now().Unix())
	startTime:=time.Now()
	c := make(chan int32, 5)
	for i := 0; i < cap(c); i++ {
		go source(c)
	}
	rnd:=<-c
	fmt.Println(time.Since(startTime))
	fmt.Println(rnd)

}
