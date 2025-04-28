// restoring sequence
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type message struct {
	msg string
	wt  chan bool
}

func generator(a string) <-chan string {
	c := make(chan string)
	waitforIt := make(chan bool)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s:%d", a, i)
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		}
	}()
	return c
}

func fanin(a, b <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-a
		}
	}()

	go func() {
		for {
			c <- <-b
		}
	}()
	return c
}

func main() {

}
