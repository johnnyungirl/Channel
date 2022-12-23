package main

import "log"

func A(name string) chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			c <- name
		}
		close(c)
	}()
	return c
}
func B(name string) chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			c <- name
		}
		close(c)
	}()
	return c
}
func FanIn(a, b chan string) chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case <-a:
				c <- <-a
			case <-b:
				c <- <-b
			}

		}
	}()
	return c
}

func main() {
	Order1 := A("coffe")
	Order2 := B("bread")
	serve:=FanIn(Order1,Order2)
	for i:=0;i<5;i++{
		log.Println(<-serve)
	}

}
