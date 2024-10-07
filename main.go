package main

import (
	"fmt"
)

type Sender interface {
	Send(msg string) error
}

type Email struct {
	Addres string
}

func (e *Email) Send(msg string) error {
	fmt.Printf("Message \"%v\" for going email %v \n", msg, e.Addres)
	return nil
}

func Notify(s Sender) {
	err := s.Send("Notify message")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Success")
}

func main() {
	email := &Email{"m82469"}
	Notify(email)
}
