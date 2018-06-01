package main
import (
    "fmt"
    "log"
    "os"
)
var (
    fileInfo *os.File
    err     error
)
func main() {
    fileInfo,err = os.Stat("test.txt")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Filename:",fileInfo.Name())
    fmt.Println("Size:",fileInfo.Size())
    fmt.Println("Permission:",fileInfo.Mode())
    fmt.Println("last modified:",fileInfo.ModTime())
    fmt.Println("Is Directory:",fileInfo.IsDir())
    fmt.Println("System Info:",fileInfo.Sys())
    

}

