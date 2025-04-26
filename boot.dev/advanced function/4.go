// anonymous function

package main

import "fmt"

func agg(add func(int, int) int, a int, b int, c int) int {
	return add(add(a, b), c)
}

func main() {
	fmt.Println(agg(func(a int, b int) int {
		return a + b
	}, 8, 4, 5))
}
