// methods in struct
package main

import "fmt"

type a struct {
	b int
	c int
}

func (rec a) check() int {
	return rec.b + rec.c
}

func main() {
	obj := a{
		b: 5,
		c: 7,
	}

	fmt.Println(obj.check())

}
