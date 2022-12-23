package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func Findfile1(c chan int, str string) {
	file1, err := ioutil.ReadFile("file1.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := strings.Count(string(file1), str)
	c <- result
}
func Findfile2(c chan int, str string) {
	file2, err := ioutil.ReadFile("file2.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := strings.Count(string(file2), str)
	c <- result
}
func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	go Findfile1(c1, "random")
	go Findfile2(c2, "Golang")
	fmt.Println(<-c2 + <-c1)

}
