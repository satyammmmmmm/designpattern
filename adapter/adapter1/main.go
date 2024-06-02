//lets take a application suppose i have bhim api and want to integrate in amazon eccommerce
//  but interface in amazon ecommerce code letscall it
// amazoninterface doesnot compatible with bhim api interface.

package main

import "fmt"

type customer struct{}     //this is customer
type bhimPay struct{}      //if customer want to pay using bhim payment option
type amazonWallet struct{} //if customer want to pay using amazonWallet payment option

// we create a adapter so that customer can have option to make payment either via bhim or amazonwallet
type bhimAdapter struct {
	paybybhim bhimPay
}

// this is bhim api interface (we want to make bhimpayment func compatible with our amazon ecommerce interface)
type bhimInterface interface {
	bhimPayment()
}

// amazoninterface is interface that is implemnted in our codebase
type amazonInterface interface {
	payViaAmazonWallet()
}

func (a amazonWallet) payViaAmazonWallet() {
	fmt.Println("payment is happening through amazon wallet")
}

func (b bhimPay) bhimPayment() {
	fmt.Println("payment is happening through bhim api ")
}

// this function help to redirect payment to bhim api
func (adapter bhimAdapter) payViaAmazonWallet() {
	fmt.Println("this is where adapter help to give option to customer to make any payment")
	adapter.paybybhim.bhimPayment()
}
func (c customer) MakePayment(pay amazonInterface) {
	pay.payViaAmazonWallet()
}
func main() {
	buyer := customer{}
	// amazonPayment := amazonWallet{}
	payBybhim := bhimAdapter{}

	buyer.MakePayment(payBybhim)

}
