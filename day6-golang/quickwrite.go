package main

import (
    "io/ioutil"
    "log"
)

func main() {
    err := ioutil.WriteFile("testquick.txt", []byte("Hi Welcome to CTS\n"), 0666)
    if err != nil {
        log.Fatal(err)
    }
}
