//select

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func logmsg(chEmails, chsms <-chan string) {
	for {
		select {
		case email, ok := <-chEmails:
			if !ok {
				fmt.Println("email channel was closed")
				return
			}
			logEmail(email)
		case sms, ok := <-chsms:
			if !ok {
				fmt.Println("sms channel was closed")
				return
			}
			logSms(sms)
		}
	}
}

func logSms(sms string) {
	fmt.Println("SMS:", sms)
}

func logEmail(email string) {
	fmt.Println("EMAIL:", email)
}

func gen_emails() <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			c <- fmt.Sprintf("Email no.->%d", i)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c)
	}()

	return c

}
func gen_sms() <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			c <- fmt.Sprintf("sms no.->%d", i)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c)
	}()
	return c

}
func main() {
	c := gen_emails()
	d := gen_sms()
	logmsg(c, d)
}
