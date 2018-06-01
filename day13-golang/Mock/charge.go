package main

import "fmt"

type VisaGateway struct {
    Name string
    Url  string
}

func (v VisaGateway) Charge() {
    fmt.Println("I am charging Visa -->")
}

type PaymentGateway interface {
    Charge()
}

func ChargeCustomer(g PaymentGateway) {
    g.Charge()
}

func main() {
    gateway := VisaGateway{   Name: "Visa",
			        Url:  "visa.com...",
			}
    ChargeCustomer(gateway)
}