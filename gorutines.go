package main

import (
	"fmt"
	"sync"
)

func PrintNumber(numbChan chan int) {
	result := 0
	for i := 0; i < 100; i++ {
		result += i
	}
	numbChan <- result
}
func PrintLetter(wg *sync.WaitGroup, letterChan chan string) {
	for i := 'A'; i < 'A'+26; i++ {
		letterChan <- string(i)
	}
	wg.Done()
}
func main() {
	numbChan := make(chan int)
	go PrintNumber(numbChan)
	fmt.Println(<-numbChan)
}
