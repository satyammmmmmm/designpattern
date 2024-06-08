package main

import "fmt"

// (e.g., email, SMS, push notifications)

type Email struct{}

type Sms struct{}

type Fanoutmsg struct{}

type Message interface {
	send()
}

func (e Email) send() {
	fmt.Println("sending message through email")
}
func (fanout Fanoutmsg) send() {
	fmt.Println("sending message to multiple user")
}

func (s Sms) send() {
	fmt.Println("sending message through Sms")
}

func FactoryMethod(messagetype string) Message {
	switch messagetype {
	case "Sms":
		return Sms{}
	case "Email":
		return Email{}
	case "Fanoutmsg":
		return Fanoutmsg{}
	default:
		fmt.Println("feature not avialble")
		return nil
	}
}

func main() {
	msgtype := "Sms"
	method := FactoryMethod(msgtype)
	method.send()

}
