// currying function that takes a function as argument and returns a funcation that has been enhanced
// handle middleware problems
package main

import "fmt"

func selfMath(arr func(int, int) int) func(int) int {
	return func(x int) int {
		return arr(x, x)
	}
}

func add(a int, b int) int {
	return a + b
}

func multiply(a int, b int) int {
	return a * b
}

func main() {
	square := selfMath(add)
	double := selfMath(multiply)

	fmt.Println(square(25))

	fmt.Println(double(4))

}
