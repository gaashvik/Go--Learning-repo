//embedded structs to acces var opf a struct directly instead of going through another struct

package main

import "fmt"

type a struct {
	a int
	b int
}
type b struct {
	c int
	d int
	a
}

func main() {
	wer := b{a: a{a: 8, b: 9}, c: 10, d: 45}
	fmt.Println(wer.b)

}
