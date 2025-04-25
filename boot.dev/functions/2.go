package main

import "fmt"

func main() {
	msg := 10
	max := 3
	fmt.Println("msg lenght is", msg)
	if msg >= max {
		fmt.Println("cannot send")
	} else {
		fmt.Println("can send")
	}
}
