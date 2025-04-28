// generator functions that return a channel as service
package main

import (
	"fmt"
	"time"
)

func generator(a string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s:%d", a, i)
			time.Sleep(time.Duration(200 * time.Millisecond))
		}
	}()
	return c
}

func main() {
	c := generator("id1")
	d := generator("id2")
	for i := 0; i < 5; i++ {
		fmt.Printf("%s\n", <-c)
		fmt.Printf("%s\n", <-d)
	}
	fmt.Println("that is enough")

}
