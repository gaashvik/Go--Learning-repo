// maps are passed by refference like slices

package main

import "fmt"

type user struct {
	name string
	num  int
}

func getfunc(names []string, nums []int) (map[string]user, error) {
	n := len(names)
	m := len(nums)
	if m != n {
		return nil, fmt.Errorf("lengths don't match")
	}
	mp := make(map[string]user)
	for i := 0; i < n; i++ {
		mp[names[i]] = user{
			name: names[i],
			num:  nums[i],
		}
	}
	return mp, nil
}
func main() {
	names := []string{"shubh", "mridul", "yashswi"}
	nums := []int{2, 3, 4}
	mp, err := getfunc(names, nums)
	if err != nil {
		fmt.Println(err)
	} else {
		for i, _ := range mp {
			fmt.Println(mp[i])
		}
	}
	_, ok := mp["shubh"]
	fmt.Println(ok)
	delete(mp, "shubh")
	_, ok1 := mp["shubh"]
	fmt.Println(ok1)

}
