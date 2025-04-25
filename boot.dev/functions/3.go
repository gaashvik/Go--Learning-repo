// scope limitation
package main

import "fmt"

func main() {
	max := 3
	fmt.Println("msg lenght is", max)
	if msg := 10; msg >= max {
		fmt.Println("cannot send")
	} else {
		fmt.Println("can send")
	}
}
