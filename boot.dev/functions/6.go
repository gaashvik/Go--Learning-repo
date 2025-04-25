// ignore return value by using _

package main

import "fmt"

func an() (x int, y int) { //named return val explicit return
	w := 5
	o := 6
	return w, o
}

func ab() (x int, y int) { // implicit named return
	x = 40
	y = 90
	return
}

func main() {
	e, o := ab()
	fmt.Println(e, o)
}
