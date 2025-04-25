package main

import "fmt"

func gethelp(msg string) (int, error) {
	cost, err := isValid(msg)
	if err != nil {
		return 0.0, err
	}
	return cost, nil
}

func isValid(msg string) (int, error) {
	if len(msg) > 10 {
		return 0.0, fmt.Errorf("it was wrong~")
	}
	return 30 * len(msg), nil
}

func main() {
	cost, err := gethelp("gcabdj")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(cost)

}
