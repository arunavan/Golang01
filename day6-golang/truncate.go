package main
import (
    "log"
    "os"
)
var (
    newFile *os.File
    err     error
)
func main() {
    err := os.Truncate("test.txt",100)
    if err != nil {
        log.Fatal(err)
    }
    
}

