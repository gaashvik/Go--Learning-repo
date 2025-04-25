// slices

package main

import "fmt"

type cost struct {
	day  int
	cost float64
}

func getcost(costs []cost) []float64 {
	max := 0
	for i := 0; i < len(costs); i++ {
		if max < costs[i].day {
			max = costs[i].day
		}
	}
	costbyDay := make([]float64, max+1)
	for i := 0; i < len(costs); i++ {
		costbyDay[costs[i].day] += costs[i].cost
	}
	return costbyDay
}

func main() {
	var obj [3]int
	for i := 0; i < len(obj); i++ {
		obj[i] = 5
		fmt.Println(obj[i])
	}
	slice := obj[:1] // one way to declare slice dynaminb memory allocation by extending the size of slice to double that of arr somewhere elese in memory

	fmt.Println(slice[0])

	slice1 := make([]int, 5, 10) // type, len, cap
	//capacity tells us the lenght of slice after which the slice will be reassigned to some new memory location that is double of current length

	for i := 0; i < 5; i++ {
		fmt.Println(slice1[i])
	}
	fmt.Println("-------------")

	slice2 := []int{2, 3, 4}
	for i := 0; i < 3; i++ {
		fmt.Println(slice2[i])
	}
	costArr := []cost{
		{0, 4.0},
		{1, 2.1},
		{1, 3.1},
		{5, 2.5},
	}
	costbyDay := getcost(costArr)

	for i := 0; i < len(costbyDay); i++ {
		fmt.Println("--", costbyDay[i])
	}

}
