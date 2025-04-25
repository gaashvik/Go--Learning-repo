package main

import "fmt"

type message struct {
	content  string
	sender   user
	receiver user
}
type user struct {
	name   string
	number int
}

func canSend(msg message) bool {
	if (msg.sender.name != "" && msg.sender.number != 0) && (msg.receiver.name != "" && msg.receiver.number != 0) {
		return true
	}
	return false
}
func main() {
	var del message = message{}
	del.sender.name = "shubh"
	del.sender.number = 8
	del.receiver.number = 2
	del.receiver.name = "nhj"
	fmt.Println(canSend(del))
}
