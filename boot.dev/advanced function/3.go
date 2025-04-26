//closures

package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(a int) int {
		sum += a
		return sum
	}
}

func main() {
	funky := adder()
	funky(1)
	funky(2)
	funky(3)
	fmt.Println(funky(4))
}
