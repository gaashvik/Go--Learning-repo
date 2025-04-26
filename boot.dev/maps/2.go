//nested maps

package main

import "fmt"

func main() {
	mp := make(map[rune]map[string]int)
	item := make(map[string]int)
	item["arunchal"] = 1
	item["arun"] = 1
	mp['a'] = item
	fmt.Println(mp['a'])
}
