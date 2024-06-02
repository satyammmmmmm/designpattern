//lets take a application suppose i have paytm api and want to integrate in amazon eccommerce
//  but interface in amazon ecommerce code letscall it
// amazoninterface doesnot compatible with paytm api interface.

package main

import "fmt"

type customer struct{}     //this is customer
type paytmPay struct{}      //if customer want to pay using paytm payment option
type amazonWallet struct{} //if customer want to pay using amazonWallet payment option

// we create a adapter so that customer can have option to make payment either via paytm or amazonwallet
type paytmAdapter struct {
	paybypaytm paytmPay
}

// this is paytm api interface (we want to make paytmpayment func compatible with our amazon ecommerce interface)
type paytmInterface interface {
	paytmPayment()
}

// amazoninterface is interface that is implemnted in our codebase
type amazonInterface interface {
	payViaAmazonWallet()
}

func (a amazonWallet) payViaAmazonWallet() {
	fmt.Println("payment is happening through amazon wallet")
}

func (b paytmPay) paytmPayment() {
	fmt.Println("payment is happening through paytm api ")
}

// this function help to redirect payment to paytm api
func (adapter paytmAdapter) payViaAmazonWallet() {
	fmt.Println("this is where adapter help to give option to customer to make any payment")
	adapter.paybypaytm.paytmPayment()
}
func (c customer) MakePayment(pay amazonInterface) {
	pay.payViaAmazonWallet()
}
func main() {
	buyer := customer{}
	// amazonPayment := amazonWallet{}
	payBypaytm := paytmAdapter{}

	buyer.MakePayment(payBypaytm)

}
