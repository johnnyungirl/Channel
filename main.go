package main

import "fmt"

func g() {
	for i := 0; i < 10; i++ {
		go fmt.Println()

	}
}
func p() {
	for i := 0; i < 10; i++ {
		go fmt.Println()

	}
}
func main(){
	
}
