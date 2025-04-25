//type assertion

package main

type one interface {
	getcost() int
}

type part1 struct {
	address string
	cost    int
}

type part2 struct {
	num  int
	cost int
}

func (p part1) getcost() int {
	return p.cost
}
func (p part2) getcost() int {
	return p.cost
}
func getExpense(intf one) {
	p, ok := intf.(part1)
	if ok {
		println("address:", p.address)
		println("cost:", p.getcost())
	} else {
		p, _ := intf.(part2)
		println("num:", p.num)
		println("cost:", p.getcost())
	}
}
func main() {
	obj1 := part1{
		address: "ghoj",
		cost:    35,
	}
	obj2 := part2{
		num:  3893,
		cost: 57,
	}
	getExpense(obj2)
}
