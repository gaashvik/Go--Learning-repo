//for loop

package main

import "fmt"

func bulk(a int) float64 {
	cost := 0.0
	for i := 0; i < a; i++ {
		cost += 1.0 + (0.01 * float64(i))
	}
	return cost
}
func bulk1(a int, thresh float) int {
	cost := 0.0
	for i := 0; ; i++ {
		cost += 1.0 + (0.01 * float64(i))
		if cost > thresh {
			return i
		}
	}
}

func main() {
	fmt.Println(bulk(10))
}
