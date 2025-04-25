//interface implementation

package main

import "fmt"

type intface interface {
	getvar() string
	getnum() int
}

type ay struct {
	name string
	num  int
}

func (obj ay) getvar() string {
	return obj.name
}

func (obj ay) getnum() int {
	return obj.num
}

func getall(intf intface) string {
	return fmt.Sprintf("name:%s num:%d", intf.getvar(), intf.getnum())
}

func main() {
	obj1 := ay{
		name: "shubh",
		num:  23,
	}
	fmt.Println(getall(obj1))

}
