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
	fmt.Printf("Message \"%v\" sent to the email %v \n", msg, e.Addres)
	return nil
}

type Phone struct {
	Number  int
	Balance int
}

func (p *Phone) Send(msg string) error {
	fmt.Printf("Message \"%v\" sent to the for number phone %v \n", msg, p.Number)
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

	phone := &Phone{Number: 7777, Balance: 10}
	Notify(phone)
}
