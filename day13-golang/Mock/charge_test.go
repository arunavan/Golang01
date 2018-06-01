package main

import (
    "fmt"
    "testing"
)

type MockGateway struct {
    Name string
    Url  string
}

func (m MockGateway) Charge() {
    fmt.Println("This is a fake gateway.  --> [no-op] <---")
    fmt.Println("Yay!  :) ")
}

func TestCharging(t *testing.T) {
    m := MockGateway{}
    ChargeCustomer(m)
}