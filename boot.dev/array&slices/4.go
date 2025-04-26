//shorthand for iteration

package main

import "fmt"

func main() {
	cost := []int{2, 4, 5, 6}
	for _, x := range cost {
		fmt.Println(x)
	}
}
