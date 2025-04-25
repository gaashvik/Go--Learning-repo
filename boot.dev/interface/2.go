// multiple interface
package main

import "fmt"

type intf1 interface {
	getsum() int
}
type intf2 interface {
	getmult() int
}

type structure1 struct {
	a int
	b int
}

func (obj1 structure1) getsum() int {
	return obj1.a + obj1.b
}
func (obj2 structure1) getmult() int {
	return obj2.a * obj2.b
}

func print(i intf1, w intf2) {
	fmt.Println(i.getsum())
	fmt.Println(w.getmult())
}

func main() {
	obj := structure1{
		a: 12,
		b: 13,
	}
	print(obj, obj)

}
