// callback functions
package main

import "fmt"

func f(a func(int, int) int, b int, c int) int {
	return a(b, c)
}
func a(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(f(a, 2, 3))
}
