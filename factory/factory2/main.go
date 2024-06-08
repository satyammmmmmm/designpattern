// we will defined an interface Paymentoption with a pay method,
// which is implemented by three structs: PayPal, Stripe, and Square, representing different payment options.
//  The PaymentFactoryUtility function creates a payment method object based on the provided option string,
//  and the main function demonstrates how to use this factory utility to make payments.

package main

import "fmt"

// Paymentoption interface defines the method pay(), which should be implemented by payment methods.
type Paymentoption interface {
	pay(int)
}

// PayPal struct represents the PayPal payment method.
type PayPal struct{}

// Stripe struct represents the Stripe payment method.
type Stripe struct{}

// Square struct represents the Square payment method.
type Square struct{}

// pay method for PayPal implementation of Paymentoption interface.
func (p PayPal) pay(amount int) {
	fmt.Printf("%d is paid through PayPal", amount)
	fmt.Println()
}

// pay method for Stripe implementation of Paymentoption interface.
func (stripe Stripe) pay(amount int) {
	fmt.Printf("%d is paid through Stripe", amount)
	fmt.Println()
}

// pay method for Square implementation of Paymentoption interface.
func (square Square) pay(amount int) {
	fmt.Printf("%d is paid through Square", amount)
	fmt.Println()
}

// PaymentFactoryUtility creates payment method objects based on the provided option string.
func PaymentFactoryUtility(option string) Paymentoption {

	switch option {
	case "PayPal":
		return PayPal{}
	case "Stripe":
		return Stripe{}
	case "Square":
		return Square{}
	default:
		fmt.Println("payment option is not supported")
		return nil
	}
}

func main() {
	// Define the payment method to be used.
	payBy := "PayPal"
	// Create payment method object using factory utility function.
	payment := PaymentFactoryUtility(payBy)
	// Check if payment method object is successfully created.
	if payment != nil {
		// Make a payment using the chosen method.
		payment.pay(100)
	}
}
