//higher order function

package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func agg(a, b, c int, add func(int, int) int) int {
	return add(add(a, b), c)
}

func main() {
	fmt.Println(agg(2, 3, 5, add))

}
